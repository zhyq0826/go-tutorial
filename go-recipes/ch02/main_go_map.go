package main

import (
	"fmt"
	"sort"
)

func main() {
	langs := map[string]string{
		"EL": "Greek",
		"ES": "Spanish",
	}

	//go 检测 key 是否存在
	_, ok := langs["en"]
	if !ok {
		fmt.Println("en is not existed")
	}

	// if 检测
	if lan, ok := langs["ES"]; ok {
		fmt.Println("es is ", lan)
	}
	// map 中删除一个 key
	delete(langs, "EL")

	// make
	chapts := make(map[int]string)
	chapts[1] = "Beginning Go"
	chapts[2] = "Go fundamentls"

	//迭代
	for k, v := range chapts {
		fmt.Println(k, v)
	}

	var keys []int
	for k := range chapts {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("key:", k, "value", chapts[k])
	}
}
