package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan int, 3)
	s := make([]string, 8)
	var num int = 0
	go func() {
		for i := 0; i < 4; i++ {
			c <- i
			v := "inner=>" + strconv.Itoa(num)
			s = append(s, v)
			num++
			fmt.Println("write to c ", i)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("reading", <-c)
		v := "outer=>" + strconv.Itoa(num)
		s = append(s, v)
		num++
	}

	fmt.Println(s)
}
