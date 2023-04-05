package go_netcat

import (
	"fmt"
	"net"
	"time"
)

var (
	defaultAddr            = ":20080"
	defaultConnDeadlineSec = uint64(60 * 5)
)

type ncHandler interface {
	Handle(conn net.Conn, logger ncLogger)
}

type ncLogger interface {
	Println(v ...any)
}

type netcat struct {
	addr     string
	deadline uint64

	logger ncLogger
}

// NewNetcat create netcat object.
func NewNetcat(opts ...NetcatOptions) *netcat {
	nc := &netcat{
		deadline: defaultConnDeadlineSec,
		addr:     defaultAddr,

		logger: newMockLogger(),
	}

	for _, opt := range opts {
		opt(nc)
	}
	return nc
}

// RunHandler start listener with income handler.
func (nc *netcat) RunHandler(handler ncHandler) error {
	listener, err := net.Listen("tcp", nc.addr)
	if err != nil {
		return err
	}

	nc.logger.Println(fmt.Sprintf("Listening on '%s'", nc.addr))
	nc.logger.Println(fmt.Sprintf("Set conn deadline: %dsec", nc.deadline))

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		if err := conn.SetDeadline(
			time.Now().Add(time.Second * time.Duration(nc.deadline)),
		); err != nil {
			return err
		}

		nc.logger.Println(fmt.Sprintf("New conn from '%s'", conn.RemoteAddr()))
		go handler.Handle(conn, nc.logger)
	}
}
