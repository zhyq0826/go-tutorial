package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var c []string
	json.Unmarshal([]byte(`["10", "11"]`), &c)
	fmt.Println(c)
}
