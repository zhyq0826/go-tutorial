package main

import "fmt"
import "sync"

type request struct {
	key   int
	value string
	ret   chan string
}
type ConcurrentMap struct {
	ch   chan request
	init sync.Once
}

func (cm *ConcurrentMap) Set(key int, value string) string {
	cm.init.Do(func() {
		cm.ch = make(chan request) //对 cm 中的 ch 进行初始化
		go runMap(cm.ch)           //传给 runMap 让其进行并发调用
	})
	result := make(chan string)
	//一个新的 request 赋值给 map 的 ch，然后等待新的 request 的 ret
	cm.ch <- request{key, value, result}
	return <-result
}

func runMap(c chan request) {
	m := make(map[int]string)
	for {
		req := <-c // c 是 map 中的 request，也是那个产生的那个新的 request
		fmt.Println(req.key, req.value)
		old := m[req.key]
		m[req.key] = req.value
		req.ret <- old
	}
}

// concurrentMap 可以再多个 goroutine 中使用而不用担心其数据的安全性，也就是 map 是 goroutine 安全的
func main() {
	var m ConcurrentMap
	fmt.Printf("Set %s\n", m.Set(1, "foo"))
	fmt.Printf("Set %s\n", m.Set(1, "bar"))
}
