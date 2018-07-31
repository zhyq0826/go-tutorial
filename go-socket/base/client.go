package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8888", 2*time.Second)
	checkErr(err)

	_, err = conn.Write([]byte("hello world"))
	checkErr(err)

	fmt.Println("read")
	var result [512]byte
	_, err = conn.Read(result[0:])
	checkErr(err)
	fmt.Println(string(result[0:]))
	os.Exit(0)
}
