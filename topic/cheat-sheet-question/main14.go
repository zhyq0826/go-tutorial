package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	//%v 对应的 struct，这里 p 是一个指针
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	fmt.Println(p.String())
}
