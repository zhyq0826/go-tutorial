package main

import (
	"fmt"
)

func main() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i < 5 {
		goto Here
	}

	// for init;condition;post
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	} // j 在循环结束之后会消失

	// undefined: j
	// fmt.Println(j)

	// for condition
	i = 0 //不能写成 i := 0，这是声明和赋值语句，因为 i 已经存在不需要再次声明
	for i < 5 {
		i++
	}

	// while true
	// for {}

	fmt.Println(i)

J:
	for j := 0; j < 5; j++ {
		for k := 0; k < 5; k++ {
			if k > 3 {
				break J
			} else {
				fmt.Println("inner for ", k)
			}
		}
	}

}
