package main

import (
	"bufio"
	"fmt"
)

type R struct{}

func (r *R) Read(p []byte) (n int, err error) {
	fmt.Println("Read")
	copy(p, "abcd") //总是有数据返回
	return 4, nil
}

func main() {
	r := new(R)
	br := bufio.NewReader(r)
	buf := make([]byte, 2)
	n, err := br.Read(buf) //第一次读取 buf size 是 2，底层 Reader 返回的是 abcd, bufio 内部的 buffer 中还剩下 cd
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read = %q\n", buf)

	buf = make([]byte, 4) //第二次调用 由于 buffer 中还有数据，不会去底层读取 abcd, 而是返回 cd
	n, err = br.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read = %q\n", buf[:n])

}
