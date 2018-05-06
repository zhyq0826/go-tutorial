package main

import (
	"fmt"
	// user go package alias
	fav "github.com/zhyq0826/go-tutorial/go-recipes/ch01/lib"
)

func main() {
	fmt.Println("default favorite packages ===>")
	fav.PrintfFavorites()
	fav.Add("github.com/onsi/ginkgo")
	fav.PrintfFavorites()
	count := len(fav.GetAll())
	fmt.Printf("Total packages in the favorites list :%d", count)
}
