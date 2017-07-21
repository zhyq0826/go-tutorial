package main

import (
	"fmt"
)

func main() {
	fmt.Println("sum1", sum1(10, 1))
	fmt.Println("sum2", sum2(10, 1))
	fmt.Println("sum3", sum3(10, 1))
	fmt.Println(sum4(1, 2))
	fmt.Println(sum5(1, 2))
	fmt.Println("deff", deff())
	fmt.Println("sum6", sum6(1, 2, 3, 4, 5))
	//函数作为值
	func_a := func() {
		fmt.Println("a")
	}
	func_a()

	//闭包
	fun_b := plus_x(1)
	fmt.Println(fun_b(2))

	var s stack
	s.push(10)
	println(s.pop())

	//斐波那契
	for _, x := range fibonacci(5) {
		println(x)
	}
}

//函数 func name(name type, name type) type {}
func sum1(x int, y int) int {
	return x + y
}

//函数 func name(name, name type) type {}
func sum2(x, y int) int {
	return x + y
}

//name return value
func sum3(x, y int) (z int) {
	z = x + y
	return
}

func sum4(x, y int) (int, int) {
	return x, y
}

func sum5(x, y int) (z int, u int) {
	z = x
	u = y
	return
}

//延迟执行的函数调用
func deff() (ret int) {
	defer func() {
		fmt.Println("1 defer")
		ret++
	}()

	defer func() {
		fmt.Println("2 defer")
		ret++
	}()

	return 0
}

//变参
func sum6(arg ...int) int {
	z := 0
	for _, n := range arg {
		z += n
	}
	return z
}

//闭包，使用 parent scope 中的变量
func plus_x(x int) func(int) int {
	return func(y int) int { return x + y }
}

//声明栈类型
type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func fibonacci(value int) []int {
	// slice
	x := make([]int, 10)
	x[0], x[1] = 1, 1
	for n := 2; n < 10; n++ {
		x[n] = x[n-1] + x[n-2]
	}

	return x
}
