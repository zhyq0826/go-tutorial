package main

import (
	"string"
	"time"
)

//目的是为了在多个 goroutine 之间可以对 map 进行原子操作

type RequestType int

const (
	Set = iota
	Get
	BeginTransaction
	EndTransaction
)

type Request struct {
	requestType RequestType //初始值是 set
	key         int
	value       string
	ret         chan string
	transaction chan Request
}

func get(m chan Request, key int) string {
	result := make(chan string)
	m <- Request{Get, key, "", result, nil}
	return <-result
}
func set(m chan Request, key int, value string) {
	m <- Request{Set, key, value, nil, nil} //
}
func beginTransaction(r chan Request) chan Request {
	t := make(chan Request)
	r <- Request{BeginTransaction, 0, "", nil, t}
	return t
}
func endTransaction(r chan Request) {
	r <- Request{EndTransaction, 0, "", nil, nil}
}

func HandleRequests(m map[int]string,
	r chan Request) {
	for {
		req := <-r
		switch req.requestType {
		case Get:
			req.ret <- m[req.key]
		case Set:
			m[req.key] = req.value
		case BeginTransaction:
			HandleRequests(m, req.transaction)
		case EndTransaction:
			return
		}
	}
}

func runMap(r chan Request) {
	//共享 map
	m := make(map[int]string)
	//开启一个事务，此时 r 的 requestType 是 beginTransaction
	beginTransaction(r)
	HandleRequests(m, r)
}

func main() {
	r := make(chan Request)
	go runMap(r)
	time.Sleep(10000000000)
}
