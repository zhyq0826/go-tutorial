package main

import (
	"fmt"
)

func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	// str2 和 str1 指向相同的一个底层数组
	str2[1] = "new"
	fmt.Println(str1)
	// str2 是新的数组
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
}
