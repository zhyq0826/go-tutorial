package main

import "fmt"

//函数可以当做参数传递，函数参数只需要声明，不需要函数体
func SplitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println(x, y)
	x, y = f(50)
	fmt.Println(x, y)
}

func main() {
	a, b := 5, 8
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x

		return x, y
	}

	SplitValues(fn)
	x, y := fn(20)
	fmt.Println(x, y)
}
