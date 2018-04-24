package main

import (
	"fmt"
)

// const 可以不需要指定类型， const 只能有 string number boolean go 会从 right side 进行推断
const i = 1

// 如果一组 const group 可以只声明第一个 const 的值，go 会自动把后续的 const 变成和第一个是 一样的
const (
	j = 0
	x
	y
	k = 1
	z
)

const u string = "abc"

// const 可以使用 constant generator 生成一个序列
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// In a const de clarat ion, the value of
// iota begins at zero and increments by one for each item in the sequence
type Flags uint

//each constant evaluates value of 1 << iota
const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func main() {
	fmt.Println("i ==> ", i)
	fmt.Println("j, x, y ==> ", j, x, y)
	fmt.Println("k, z ==> ", k, z)
	fmt.Println("u ==> ", u)

	fmt.Println(" weekday ==> ", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(" flags ==> ", FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}
