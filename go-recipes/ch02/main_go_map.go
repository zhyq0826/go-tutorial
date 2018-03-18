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

	_, ok := langs["en"]
	if !ok {
		fmt.Println("en is not existed")
	}

	if lan, ok := langs["ES"]; ok {
		fmt.Println("es is ", lan)
	}
	delete(langs, "EL")

	chapts := make(map[int]string)
	chapts[1] = "Beginning Go"
	chapts[2] = "Go fundamentls"

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
