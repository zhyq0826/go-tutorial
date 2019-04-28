package main

import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch s := v.(type) {
	case *student, student:
		fmt.Printf("%T", s)
		fmt.Printf("%V", s)
		fmt.Println(s.Name)
	}
}

func main() {
	zhoujielun(student{"string"})
}
