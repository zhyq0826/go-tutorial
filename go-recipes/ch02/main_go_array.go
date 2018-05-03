package main

import "fmt"

func main() {
	// 空数组
	var x [5]int
	// 字面量空数组
	y := [5]int{1, 2, 3, 4, 5}
	// 可以声明字面量的位置，其他的缺省
	langs := [4]string{0: "Go", 3: "Julia"}
	langs[1] = "Rust"
	// 多行隔开，最后一个 coma 是必须的
	z := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	//变长数组
	k := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%#v\n", langs)
	fmt.Printf("%#v\n", z)
	fmt.Printf("%#v\n", k)
	// for 循环
	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]:%s\n", i, langs[i])
	}

	// range 返回 k 和 v k 是索引
	for k, v := range langs {
		fmt.Printf("langs[%d]:%s\n", k, v)
	}

}
