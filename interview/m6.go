package main

import (
	"fmt"
)

/*
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4

A deferred function's arguments are evaluated when the defer statement is evaluated.

deferred 的函数参数再 defer 时已经开始计算了，所以 calc("10", a, b) 先开始计算
*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
