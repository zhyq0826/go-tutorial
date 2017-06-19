package main

import (
	"fmt"
)

func main() {
	var arr [10]int //默认初始化为0值
	fmt.Printf("%T", arr)
	arr[0] = 1
	arr[1] = 2
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	a := [3]int{1, 2, 3} //复合声明
	for i := 0; i < len(a); i++ {
		println(a[i])
	}

	b := [...]int{1, 2, 3} //复合声明
	for i := 0; i < len(b); i++ {
		println(b[i])
	}

	//invalid type for composite literal: int 数组声明 必须是 []int
	c := [3][2]int{[...]int{1, 2}, [...]int{1, 2}, [...]int{1, 2}}
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			println(c[i][j])
		}
	}

	d := [...]int{
		1,
		2,
		3,
		4, // comma 是必须的
	}

	for _, value := range d {
		println(value)
	}

	//var arr [5]int{1, 2, 3, 4, 5} 没有这种声明和初始化方式
	e := [5]int{1, 2, 3, 4, 5}
	for _, value := range e {
		println(value)
	}

}
