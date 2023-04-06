package go_netcat

import (
	"bytes"
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
func (h *handlerExec) Handle(conn net.Conn) {
	rp, wp := io.Pipe()

	h.cmd.Stdin = conn
	h.cmd.Stdout = wp

	go func() {
		_, _ = io.Copy(conn, rp)
	}()
	if err := h.cmd.Run(); err != nil {
		_, _ = io.Copy(conn, bytes.NewReader([]byte(err.Error())))
	}
	defer func() {
		conn.Close()
	}()
}
