package main

import (
	"fmt"
)

func main() {

	// 使用 make 创建 slice
	sl := make([]int, 5) //创建一个 length 为 5 的 slice，slice 的底层总是指向一个 array
	fmt.Println(len(sl))
	fmt.Println(cap(sl))

	var arr [10]int
	slice := arr[0:5]
	fmt.Println(cap(slice)) //cap 指向的实际的底层数组的长度 10
	fmt.Println(len(slice)) //len 只表示当前类型的长度 5
	// slice[11] = 10 runtime error: index out of range slice 实际指向底层的 array

	// append 创建新的 slice并返回，如果原来底层的 array 容量够用，底层的 array 并不会改变
	s1 := append(slice, 5)
	s2 := append(slice, 5, 5, 5)
	// s3 := append(slice, arry...)
	fmt.Println(s1)
	fmt.Println(s2)
	// print(s3)

	s3 := []int{1, 2, 3}
	s4 := make([]int, 2)
	copy(s4, s3) //copy from s3
	fmt.Println(s4)

}
