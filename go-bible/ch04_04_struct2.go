package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	// pp := &Point{1, 2}
	pp := new(Point)
	*pp = Point{1, 2}
	fmt.Printf("%#v", *pp)
}
