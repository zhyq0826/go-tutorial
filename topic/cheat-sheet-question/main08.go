package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	//指针不同
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	//struct 是 可以比较的 filed  相同
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
