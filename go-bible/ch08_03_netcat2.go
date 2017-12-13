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
	mustCopy(conn, os.Stdin)
}

//从标准输入中读取数据并且送到 conn
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
