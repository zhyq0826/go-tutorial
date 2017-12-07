package main

import (
	"fmt"
)

func main() {
	var x1 interface{}
	var x2 int
	var x3 string
	var x4 bool

	fmt.Println(sqlQuote(x1))
	fmt.Println(sqlQuote(x2))
	fmt.Println(sqlQuote(x3))
	fmt.Println(sqlQuote(x4))
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "null"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "true"
		}

		return "false"
	case string:
		return "string"
	default:
		panic("no match case")
	}
}
