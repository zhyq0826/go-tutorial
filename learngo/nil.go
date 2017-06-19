package main

import (
	"fmt"
)

// %v  the value in a default format
//     when printing structs, the plus flag (%+v) adds field names
// %#v a Go-syntax representation of the value
// %T  a Go-syntax representation of the type of the value
// %%  a literal percent sign; consumes no value
func main() {
	fmt.Printf("%T", nil)
	fmt.Println(nil)
	fmt.Printf("%#v", nil)
	// println(nil)
}
