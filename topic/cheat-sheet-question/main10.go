package main

import (
	"fmt"
)

type Stduent struct {
	Age int
}

func main() {
	kv := map[string]Stduent{"menglu": {Age: 21}}
	// kv["menglu"].Age = 22 //cannot assign to struct field kv["menglu"].Age in map
	s := []Stduent{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}
