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
			num++
			v := "inner=>" + strconv.Itoa(num)
			s = append(s, v)
			fmt.Println("write to c ", i)
		}
	}()

	for i := 0; i < 4; i++ {
		num++
		v := "outer=>" + strconv.Itoa(num)
		s = append(s, v)
		fmt.Println("reading", <-c)
	}

	fmt.Println(s)
}
