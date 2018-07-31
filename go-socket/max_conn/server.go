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
	var i int
	for {
		time.Sleep(time.Second * 10)
		if _, err := l.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		}
		i++
		log.Printf("%d: accept a new connection\n", i)
	}
}
