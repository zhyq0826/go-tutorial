package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Sum(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}

	return total
}

func main() {
	x, y := 20, 10
	result := Add(x, y)
	fmt.Println("[Add]:", result)
	result = Subtract(x, y)
	fmt.Println("[Subtract]:", result)
	a, b := "shiju", "Varghere"
	a, b = Swap(a, b)
	fmt.Println("swap a and b", a, b)
	fmt.Println("sum x and y", Sum(x, y))
	s := []int{1, 2, 3, 4, 5}
	// sum slice dot is after
	fmt.Println("sum slice", Sum(s...))
}
