package main

import (
	"fmt"
)

type Request struct {
	key   int
	value string
	ret   chan string
}

// 创建 request，并且返回 request 的 result 的值
func set(m chan Request, key int, value string) string {
	result := make(chan string)
	m <- Request{key, value, result}
	return <-result
}

func runMap(c chan Request) {
	// map 是本地的，通过穿 request 来达到共享 map 数据的作用
	m := make(map[int]string)
	for {
		req := <-c
		//获取本地 map req.key 的值，把 req 的新值赋给本地，返回旧值
		old := m[req.key]
		m[req.key] = req.value
		req.ret <- old
	}
}

func main() {
	req := make(chan Request)
	go runMap(req)
	fmt.Println("old value is ", set(req, 1, "foo"))
	fmt.Println("old value is ", set(req, 1, "bar"))
	fmt.Println("old value is ", set(req, 1, "foo"))
}
