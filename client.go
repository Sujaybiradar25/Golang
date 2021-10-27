package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("Couldn't establish connection")
		conn.Close()
		return
	}
	message := os.Args
	fmt.Println("Starting connection with IP ", conn.LocalAddr())
	fmt.Fprintf(conn, "%s", message[1])
	fmt.Println("Sent message", message[1])
	return
}
