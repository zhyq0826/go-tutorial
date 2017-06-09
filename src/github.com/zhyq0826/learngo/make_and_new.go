package main

import (
	"fmt"
)

func main() {
	var ptr1 *int             //初始化为 0 值即 nil 的指针
	fmt.Printf("%v\n", ptr1)  // <nil>
	fmt.Printf("%+v\n", ptr1) // <nil>

	var ptr2 *int
	a := 0
	ptr2 = &a
	fmt.Printf("%v\n", ptr2)   // a 的 memory address 0xc4200721b0
	fmt.Printf("%+v\n", ptr2)  // 0xc4200721b0
	fmt.Printf("%+v\n", *ptr2) // memory address 为 0xc4200721b0 的位置存放的数据
	*ptr2 = 3                  //通过 ptr2 改变 a 位置的值
	fmt.Printf("%v\n", ptr2)   //ptr2 的值不变 0xc4200721b0
	fmt.Printf("%v\n", a)      //a 改变
	fmt.Printf("%v\n", *ptr2)  //ptr2 指向的内容改变

	var m1 map[int]string // 声明一个空 map, 未初始化只能说明其类型是 map[int]string，但其值是个 nil
	fmt.Printf("%v\n", m1)
	// m1[10] = "a" map[]panic: assignment to entry in nil map
	// map 通过 make 初始化或者通过字面量
	m2 := make(map[int]string)
	m2[10] = "a"
	fmt.Printf("%v\n", nil)
	// fmt.Printf("%v\n", &nil)
	// fmt.Printf(nil == nil) nil 是不能比较的
}
