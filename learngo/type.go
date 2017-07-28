package main

import (
	"fmt"
)

type Person map[string]string
type Cat map[string]string

func keys(m map[string]string) (keys []string) {
	for key, _ := range m {
		keys = append(keys, key)
	}

	return
}

func name(p Person) string {
	return p["name"]
}

func main() {

	var person = Person{"name": "Rob"}
	var cat = Cat{"name": "June"}

	fmt.Printf("%v\n", name(person))
	//cannot use cat (type Cat) as type Person in argument to name
	//fmt.Printf("%v", name(cat))

	fmt.Printf("%v\n", keys(person))
	fmt.Printf("%v\n", keys(cat))
}
