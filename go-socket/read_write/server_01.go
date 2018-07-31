package main

import (
	"fmt"
	"log"
	"net"
)

func handleConn(c net.Conn) {
	// defer c.Close()
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
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
