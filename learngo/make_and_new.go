package main

import (
	"fmt"
)

func main() {
	var ptr1 *int             //初始化为 0 值即 nil 的指针
	fmt.Printf("%#v\n", ptr1) // (*int)(nil)
	fmt.Printf("%p\n", ptr1)  // nil 的地址 ==>0x0

	var ptr *string          //初始化为 0 值即 nil 的指针
	fmt.Printf("%#v\n", ptr) // (*string)(nil)
	fmt.Printf("%p\n", ptr)  // nil 的地址 ==> 0x0

	var ptr2 *int
	a := 0
	ptr2 = &a
	fmt.Printf("%#v\n", ptr2)  // (*int)(0xc4200721b0)
	fmt.Printf("%#v\n", *ptr2) // memory address 为 0xc4200721b0 的位置存放的数据 ==> 0
	*ptr2 = 3                  //通过 ptr2 改变 a 位置的值
	fmt.Printf("%#v\n", ptr2)  //ptr2 的值不变 ==> (*int)0xc4200721b0
	fmt.Printf("%#v\n", a)     //a 改变 ==> 3
	fmt.Printf("%#v\n", *ptr2) //ptr2 指向的内容改变 ==>3

	var m1 map[int]string   // 声明一个空 map, 未初始化只能说明其类型是 map[int]string，但其值是个 nil
	fmt.Printf("%#v\n", m1) //map[int]string(nil)
	fmt.Println(m1 == nil)  //true
	fmt.Printf("%p\n", m1)  //0x0
	// m1[10] = "a" map[]panic: assignment to entry in nil map

	// map 通过 make 初始化或者通过字面量
	m2 := make(map[int]string)
	fmt.Println(m2 == nil) //false
	m2[10] = "a"
	fmt.Printf("%#v\n", m2) //map[int]string{10:"a"}

	// new 返回 map 的指针
	ptr_m3 := new(map[int]string)
	fmt.Printf("%#v\n", ptr_m3)  //&map[int]string(nil)
	ptr_m3 = &m2                 //ptr_m3 是一个指针，可以指向同类型的 m2
	fmt.Printf("%#v\n", ptr_m3)  //&map[int]string{10:"a"}
	fmt.Printf("%#v\n", *ptr_m3) //直接输出 ptr_m3 指向的 m2 ==> map[int]string{10:"a"}
	(*ptr_m3)[11] = "b"          //
	fmt.Printf("%#v\n", m2)      //map[int]string{10:"a", 11:"b"}

	// cannot take the address of nil, use of untyped nil
	// fmt.Printf("%v\n", &nil)
	//nil 是不能比较的
	// fmt.Printf(nil == nil)
}
