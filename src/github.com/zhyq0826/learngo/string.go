package main

import (
	"fmt"
)

func main() {
	var ss = "I am string"
	fmt.Println(ss)

	sss := "hello"
	ccc := []rune(sss)
	ccc[0] = 'c'
	fmt.Println(string(ccc))

	ms := "start" +
		"end"

	fmt.Println(ms)

	msm := `start
        end`
	fmt.Println(msm)
}
