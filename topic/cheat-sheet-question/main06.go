package main

import (
	"encoding/json"
	"fmt"
)

type Girl struct {
	Name       string `json:"name"`
	DressColor string `json:"dress_color"`
}

func (g Girl) SetColor(color string) {
	g.DressColor = color
}
func (g Girl) JSON() string {
	data, _ := json.Marshal(&g)
	return string(data)
}
func main() {
	g := Girl{Name: "menglu"}
	g.SetColor("white")
	fmt.Println(g.JSON())
}
