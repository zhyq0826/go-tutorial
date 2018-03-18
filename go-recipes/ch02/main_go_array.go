package main

import "fmt"

func main() {
	var x [5]int
	y := [5]int{1, 2, 3, 4, 5}
	langs := [4]string{0: "Go", 3: "Julia"}
	langs[1] = "Rust"
	z := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	k := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%#v\n", langs)
	fmt.Printf("%#v\n", z)
	fmt.Printf("%#v\n", k)

	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]:%s\n", i, langs[i])
	}

	for k, v := range langs {
		fmt.Printf("langs[%d]:%s\n", k, v)
	}

}
