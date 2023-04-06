package main

import (
	"log"

	go_netcat "github.com/mr-chelyshkin/go-netcat"
)

func main() {
	l := log.Default()
	l.Println()

	h := go_netcat.NewHandlerExec()

	nc := go_netcat.NewNetcat(
		go_netcat.WithLogger(l),
	)
	log.Fatalln(nc.RunHandler(h))
}
