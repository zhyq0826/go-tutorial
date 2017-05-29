package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

// methods
func (c *Circle) circleArea() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	fmt.Println("circle")
	var circle Circle     //initialization zero field
	changeCircleR(circle) //can't change circle value, arugment pass passed by value
	fmt.Printf("%T", circle)
	circle.r = 10
	fmt.Println(circle.x)
	fmt.Println(circle.y)
	fmt.Println(circle.r)

	fmt.Println("circle new")
	c := new(Circle) //initialization return pointer *Circle
	//cannot use c (type *Circle) as type Circle in argument to changeCircleR
	// changeCircleR(c)
	fmt.Printf("%T", c)
	fmt.Println(c.x)
	fmt.Println(c.y)
	fmt.Println(c.r)

	fmt.Println("circle literal")
	c1 := Circle{x: 1, y: 2, r: 3}
	fmt.Printf("%T", c1)
	changeCircleR(circle)
	fmt.Println(c1.x)
	fmt.Println(c1.y)
	fmt.Println(c1.r)

	c2 := new(Circle) //*Circle
	c2.x = 1
	c2.y = 2
	c2.r = 3
	fmt.Println(c2.x)
	fmt.Println(c2.y)
	fmt.Println(c2.r)
	fmt.Println(c2.circleArea())
}

func changeCircleR(c Circle) {
	c.r = 10
}
