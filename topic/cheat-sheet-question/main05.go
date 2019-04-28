package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Println("fn => ", i)
		ch <- vs[i](name)
	}
	for i := range vs {
		go fn(i)
	}
	// 一个 goroutine 执行完成之后就返回了，其他的 goroutine 得不到执行
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
