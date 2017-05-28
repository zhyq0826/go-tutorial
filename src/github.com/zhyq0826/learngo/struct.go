package main

import (
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	var circle Circle //initialization zero field
	fmt.Println(circle.r)
	fmt.Println(circle.x)
	fmt.Println(circle.y)

	c := new(Circle) //initialization
	fmt.Println(c.r)
	fmt.Println(c.x)
	fmt.Println(c.y)
}
