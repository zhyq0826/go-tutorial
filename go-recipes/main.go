package main

import (
	"fmt"
	"go-tutorial/go-recipes/ch01/lib"
	"go-tutorial/go-recipes/ch01/strutils"
)

func main() {
	fmt.Println(strutils.ToUpperCase("aaa"))
	fmt.Println(strutils.ToLowerCase("BBB"))
	fmt.Println(strutils.ToFirstUpper("ccc"))
	fmt.Println("default favorite packages ===>")
	lib.PrintfFavorites()
	lib.Add("github.com/onsi/ginkgo")
	fmt.Println("add favorite packages ===>")
	lib.PrintfFavorites()
	count := len(lib.GetAll())
	fmt.Printf("Total packages in the favorites list :%d", count)
}
