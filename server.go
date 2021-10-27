package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("error staring server")
		l.Close()
		return
	}
	defer l.Close()
	handshake, err := l.Accept()
	if err != nil {
		fmt.Println("Couldn't establish connectiom")
		l.Close()
		return
	}
	fmt.Println("Started server with ", handshake.LocalAddr())
	var buf bytes.Buffer

	io.Copy(&buf, handshake)

	if err == nil {
		fmt.Println("Message received : ", buf.String())
	}
}
