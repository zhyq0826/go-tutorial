package main

import (
	"fmt"
)

type myFuncType func(int, int) bool

// Within a list of parameters or results,
// the names (IdentifierList) must either all be present or all be absent.
// If present, each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique.
// If absent, each type stands for one item of that type.
// functype 中 parameter 要么全是带名字的，要么都不是带名字的，如果不是带名字的，一个 type 仅表示一个参数，名称必须唯一
// type myFuncType2 func(x int, int) bool ==> invalid

// Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.
// 参数和结果集必须要带()，除非结果集只有一个 unnamed type
// type myFuncType3 func(x int, y int) z bool == invalid
type myFuncType2 func(int, int) (z bool)
type myFuncType3 func(x int, y int) bool

func createMyFunc(myFunc myFuncType) {
	fmt.Println("createMyFunc")
}

func createMyFunc2(myfunc myFuncType2) {
	fmt.Println("createMyFunc2")
}

func createMyFunc3(myfunc myFuncType3) {
	fmt.Println("createMyFunc3")
}

func main() {
	createMyFunc(myFuncType(addFunc))
	createMyFunc(myFuncType(addFunc2))
	createMyFunc(myFuncType(addFunc3))
	createMyFunc(myFuncType(addFunc4))
	createMyFunc(myFuncType(addFunc5))

	createMyFunc2(myFuncType2(addFunc))
	createMyFunc2(myFuncType2(addFunc2))
	createMyFunc2(myFuncType2(addFunc3))
	createMyFunc2(myFuncType2(addFunc4))
	createMyFunc2(myFuncType2(addFunc5))

	createMyFunc3(myFuncType3(addFunc))
	createMyFunc3(myFuncType3(addFunc2))
	createMyFunc3(myFuncType3(addFunc3))
	createMyFunc3(myFuncType3(addFunc4))
	createMyFunc3(myFuncType3(addFunc5))
}

func addFunc(x, y int) bool {
	return x == y
}

func addFunc2(x int, y int) bool {
	return x == y
}

func addFunc3(x, y int) (z bool) {
	return x == y
}

func addFunc4(x, y int) (z bool) {
	z = x == y
	return z
}

func addFunc5(x int, y int) (z bool) {
	z = x == y
	return z
}
