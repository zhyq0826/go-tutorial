package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("%#v\n", f)
	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, " is string", vv)
		case float64:
			fmt.Println(k, " is float64", vv)
		case []interface{}:
			fmt.Println(k, " is array", vv)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println("don't known k")
		}
	}
}
