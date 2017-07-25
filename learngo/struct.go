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
	fmt.Println("var circle c1")
	var c1 Circle     //initialization zero field
	changeCircleR(c1) //can't change circle value, arugment pass passed by value
	fmt.Printf("%T\n", c1)
	c1.r = 10 //虽然 c1 是个值，但是是 struct 的实例类型，c1.r 其实是 &c1.r 会改变 c1 的值
	fmt.Println("c1 x ", c1.x)
	fmt.Println("c1 y ", c1.y)
	fmt.Println("c1 r ", c1.r)

	fmt.Println("new circle c2")
	c2 := new(Circle) //initialization return pointer *Circle
	//cannot use c (type *Circle) as type Circle in argument to changeCircleR
	// changeCircleR(c)
	fmt.Printf("%T\n", c2)
	fmt.Println("c2 x ", c2.x) //虽然 c2 是指针类型，在调用时候并不需要 &c 这样的写法，go 会自动进行转换
	fmt.Println("c2 y ", c2.y)
	fmt.Println("c2 r ", c2.r)

	fmt.Println("circle literal")
	c3 := Circle{x: 1, y: 2, r: 3}
	fmt.Printf("%T\n", c3)
	changeCircleR(c3)
	fmt.Println("c3 x ", c3.x)
	fmt.Println("c3 y ", c3.y)
	fmt.Println("c3 r ", c3.r)

	c4 := new(Circle) //*Circle
	c4.x = 1
	c4.y = 2
	c4.r = 3
	fmt.Println("c4 x ", c4.x)
	fmt.Println("c4 y ", c4.y)
	fmt.Println("c4 r ", c4.r)
	fmt.Println("c4 area", c4.circleArea())
}

func changeCircleR(c Circle) {
	c.r = 10
}
