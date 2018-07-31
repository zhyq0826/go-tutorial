package main

import (
	"bytes"
	"fmt"
	"net"
)

func handleConn(c net.Conn) {
	// defer c.Close()
	var result [512]byte
	c.Read(result[0:])
	c.Write(bytes.ToLower(result[0:]))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listening err", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accpet err", err)
			break
		}

		fmt.Println(c.RemoteAddr())
		go handleConn(c)
	}
}
