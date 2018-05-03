package main

import (
	"fmt"
)

func main() {
	var x []int
	// slice 可以和 nil 比较
	if x == nil {
		fmt.Println("x is nil")
	}

	y := make([]int, 3, 5)
	y[0] = 10
	y[1] = 20
	fmt.Println("y is ", y)

	z := []int{10, 20, 30}
	k := []int{0: 10, 2: 30}
	fmt.Println("z is ", z)
	fmt.Println("k is ", k)
	h := []int{}
	if h == nil {
		fmt.Println("h is nil")
	}

	y = append(y, 10, 10, 10)
	fmt.Println("y is ", y)
	var q []int
	// append nil slice
	q = append(q, 1)
	fmt.Println("q is ", q)
}
