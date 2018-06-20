package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// typeOf 总是返回的是一个 实际类型是 reflect.Type concrete type
	t := reflect.TypeOf(3)
	//下面两个输出是一样的
	fmt.Println(t.String()) // int
	fmt.Println(t)          //int
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File 反应的是 os.Stdout 的 concrete type

	// valueOf 总是返回 interface 的 dynamic value 实际类型是 refect.Value
	// reflect.Value 能够接受任何值
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Println(v.String()) //<int Value>
	fmt.Println(v.Type())   // int
	x := v.Interface()
	i := x.(int)
	fmt.Println(i) //3

	fmt.Printf("%T\n", reflect.TypeOf(3)) // int
	fmt.Printf("%#v\n", reflect.TypeOf(3))

	// reflect.Value 和 interface{} 都能接受任意值，差别在于 interface 因此了实现细节和 value 的固有属性，
	// 除非我们知道 dynamic type 是多少，并且使用 type assertion 获取到其值，否则什么也做不了
	// 相反使用 reflect.Value 我们可以有很多方式获取 value 的内容
}
