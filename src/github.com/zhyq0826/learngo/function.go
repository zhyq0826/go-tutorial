package main

func main() {
	println(sum_two(10, 1))
	x, y := sum_three(1, 2, 3)
	println(x, y)
	println(sum(1, 1))
	println(return_name(1, 1))
	println(deff())
	func_arg(1, 2, 3, 4, 5)

	//函数作为值
	func_a := func() {
		println("a")
	}

	func_a()

	fun_b := plus_x(1)
	println(fun_b(2))

	var s stack
	s.push(10)
	println(s.pop())

	//斐波那契
	for _, x := range fibonacci(5) {
		println(x)
	}
}

//只有一个返回值
func sum_two(x, y int) int {
	return x + y
}

//多个返回值
func sum_three(x int, y int, z int) (int, int) {
	return x + y, y + z
}

func sum(x, y int) int {
	return 0 // return 是必须的
}

//带名称的返回值
func return_name(x, y int) (z int) {
	z = x + y
	return
}

//延迟执行的函数调用
func deff() (ret int) {
	defer func() {
		ret++
	}()

	return 0
}

//变参
func func_arg(arg ...int) {
	for _, n := range arg {
		print(n)
	}
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
