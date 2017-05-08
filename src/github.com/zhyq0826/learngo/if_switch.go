package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if true {
		fmt.Println("1")
	} else {
		fmt.Println("not 1")
	}

	marks := 90
	var grade string
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70, 60, 50: // , 隔开的列表
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("You got A")
	case grade == "B":
		fmt.Println("You got B")
	default:
		fmt.Println("You got other")
	}

	//接受初始化语句
	if a := rand.Intn(100); a > 5 {
		fmt.Println(a)
	} else if a < 5 {
		fmt.Println(a)
	}

	switch {
	case grade == "A":
		fmt.Println("A")
		fallthrough //匹配成功之后继续向下执行
	case grade == "B":
		fmt.Println("B")
	}

}
