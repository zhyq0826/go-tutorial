package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
	// 连接不断读取数据并转化
	input := bufio.NewScanner(conn)
	defer conn.Close()
	for input.Scan() {
		// goroutine 的函数的参数是在执行 go 语句时计算的， 所以 input.Text() 取时当前 conn 的数据
		go echo(conn, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
