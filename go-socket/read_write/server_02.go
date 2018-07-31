package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

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

//go-tcpsock/read_write/server4.go
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		time.Sleep(10 * time.Second)
		var buf = make([]byte, 65536)
		log.Println("start to read from conn")
		c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes,  error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
