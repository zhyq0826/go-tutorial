package main

import (
	"fmt"
)

func main() {
	// 字面量定义
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for _, v := range m {
		fmt.Println(v)
	}

	v, ok := m["a"]
	fmt.Println(v)
	fmt.Println(ok) //true

	delete(m, "c")
	//不存在的值返回类型的 0 值
	v1, ok1 := m["c"]
	fmt.Println(v1)  // 0
	fmt.Println(ok1) // false
	//make 定义
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 1
	fmt.Println(a)

	if _, ok := m["d"]; ok {
		fmt.Println("d exists")
	} else {
		fmt.Println("d not exists")
	}
}
