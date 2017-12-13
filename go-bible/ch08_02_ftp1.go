package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	// "time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		out, err := exec.Command(input.Text()).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(conn, "\n", string(out))
	}
}
