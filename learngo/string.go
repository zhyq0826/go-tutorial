package main

import (
	"fmt"
)

func main() {
	var ss = "I am string"
	fmt.Println(ss)

	//single char ANSI
	fmt.Println(ss[0])

	sss := "hello"
	ccc := []rune(sss)
	ccc[0] = 'c'
	fmt.Println(string(ccc))

	// plus operator
	ms := "start" +
		"end"

	fmt.Println(ms)

	// multi line string
	msm := `start
        end`
	fmt.Println(msm)
}
