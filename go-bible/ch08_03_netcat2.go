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
	// 从 conn 中读取数据并且送到标准输出
	go mustCopy(os.Stdout, conn)
	//从标准输入中读取数据并且送到 conn
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF { //check eof ctrl + d
			os.Exit(1)
		}
	}
}
