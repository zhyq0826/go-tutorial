package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	done := make(chan struct{})
	// 从 conn 中读取数据并且送到标准输出
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	tcpConn, _ := conn.(*net.TCPConn)
	tcpConn.CloseWrite()
	<-done //wait for background goroutine to finish
}

//从标准输入中读取数据并且送到 conn
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF { //check eof ctrl + d
			return
		}
		log.Fatal(err)
	}
}
