package go_netcat

import (
	"fmt"
	"io"
	"net"
	"os/exec"
)

type handlerExec struct {
	cmd *exec.Cmd
}

// NewHandlerExec return new handlerExec object.
func NewHandlerExec() *handlerExec {
	return &handlerExec{
		cmd: exec.Command("/bin/sh", "-i"),
	}
}

// Handle exec.
func (h *handlerExec) Handle(conn net.Conn, _ ncLogger) error {
	rp, wp := io.Pipe()

	h.cmd.Stdin = conn
	h.cmd.Stdout = wp

	var errCh = make(chan error)

	go func() {
		_, err := io.Copy(conn, rp)
		if err != nil {
			errCh <- err
		}
		errCh <- fmt.Errorf("FFFFFFFF")
	}()
	if err := h.cmd.Run(); err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()

	return fmt.Errorf("asd")
}
