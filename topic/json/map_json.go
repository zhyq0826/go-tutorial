package main

import (
	"encoding/json"
	"fmt"
)


type Data struct{}

func (*Data) Read(p []byte) (n int, err error) {
	p = append(p, `{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`...)
	return cap(p), nil
}

func main() {
	dataSource := make(map[string]interface{})
	err := json.NewDecoder(new(Data)).Decode(dataSource)
	fmt.Println(err)
	fmt.Println(dataSource)
}
