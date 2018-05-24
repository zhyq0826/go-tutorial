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
type myFuncType2 func(int int) (z bool)
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
	//返回值不带名字的能够转换带名字或不带名字的
	createMyFunc(myFuncType(addFunc))
	createMyFunc(myFuncType(addFunc2))
	createMyFunc(myFuncType(addFunc3))
	createMyFunc(myFuncType(addFunc4))

	//返回值带名字的都不能转换
	// createMyFunc2(myFuncType2(addFunc)) // cannot convert addFunc (type func(int, int) bool) to type myFuncType2 addFunc 和 myFuncType2 函数签名不一致
	// createMyFunc2(myFuncType2(addFunc2)) //cannot convert addFunc2 (type func(int, int) bool) to type myFuncType2
	// createMyFunc2(myFuncType2(addFunc3)) //cannot convert addFunc3 (type func(int, int) bool) to type myFuncType2
	// createMyFunc2(myFuncType2(addFunc4)) //cannot convert addFunc4 (type func(int, int) bool) to type myFuncType2

	//返回值不带名字的能够转换带名字或不带名字的
	createMyFunc3(myFuncType3(addFunc))
	createMyFunc3(myFuncType3(addFunc2))
	createMyFunc3(myFuncType3(addFunc3))
	createMyFunc3(myFuncType3(addFunc4))
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
