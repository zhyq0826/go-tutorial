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
