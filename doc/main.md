## datatype

go types can be classifed as four category

- boolean types
- numeric types
- string types
- derived types
    - pointer
    - array
    - structure
    - union
    - function
    - slice
    - interface
    - map
    - channel

**interger type**

- unit8
- unit16
- unit32
- unit64
- int8
- int16
- int32
- int64

**float type**

- float32
- float64
- complex64
- complex128

**numeric type**

- byte uint8
- rune uint32
- unit
- int
- uintptr


## variable 

go 是类型置后的。

可选的 `;` 一般来说如果一行只有一条语句 `;` 是不需要的，如果要在一行出现多条语句 `;` 是必须的。

go 可以动态进行类型推断，如果有初始化值可以不必一开始就声明变量类型。

`:=` 赋值只能用在 function 中，表示声明并初始化。

**declare variable**

var variable_list optional_data_type


## constant

常量在 go 中只能是 boolean、string、numerical。

常量可以使用枚举、指定初始值、或者显式指定类型和初始值。


## operator

go 中只能同类型的进行操作，不能混用，也就是不能用 int32 + int16


两个重要的操作符是 &、* ，同 c。

```
var ptr *int  //声明指针
var a = 0

ptr = &a  //指针赋值
```

## string

go 中字符串是不可变的。

go 中的字符串不同于 c，不是字符的序列，而是一个个 byte，直接输出 s[0] 是 byte 代表的 10 进制数

go 中有单独的字符序列是 rune 数组。

多行字符串注意 + 号的位置。

## scope

go 是 block scope language


## slice 

slice 是 array 的一段，唯一不同的是可变。

创建 slice 可以用 make，array[start:end]。

append 的返回值必须要接住，否则报：append(class, t) evaluated but not used 


## if switch

go 的控制结构少，有 if、for switch。

if、for 接受初始化语句。

switch 不同于其他语言，不需要 break，因为匹配失败之后不会继续向下执行，可以用 fallthrough。

if 的大括号必须在同一行。


## goto for

go 有 goto 语句，标签名是大小写敏感的。

go 没有 while 语句，可以使用 for 替代。

for 有 break 和 continue 语句，break 可以指定标签来跳出哪个循环

## range 

range 对 array、slice、map 的用法

## array

声明 array 必须要有个数说明。

slice 是一个指向底层 array 的类型，即使从新的 slice 创建 slice，两个 slice 也是指向同一个 array

append() 可以向 slice 追加元素并且自动扩容，扩容之后会返回新的底层 array，所以调用 append 可能会产生新的底层 array。


## function

函数的定义
```
func (p mytype) functioname(q int) (r,s int) {return 0, 0}
```

socpe 遵循局部覆盖全部的原则。

如果函数有返回值，那 return 语句是必须的。

defer 可以让函数的执行延迟，记住必须是函数的执行，匿名函数屁股后面一定是调用符号()。

defer 表示在函数执行完成之后在执行，所以 defer 可以用来恢复异常。

函数可接受变参，参数类型一致的参数列表，其实是一个 slice。

函数可以作为值来使用，可以把函数赋值给一个变量。

## package

go package 是 go 源文件的 package 名称，和 源文件所在 directory 其实是没有关系的，不过一般意义上 path 的最后一个元素和源文件的 package 
名称相同。

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand


像 Python 一样，导入可单行，也可分组。

package 默认只导出首字母大写的变量、函数，小写不导出，也就是可以使用那些头字母是大写的函数。

包可以有别名，别名的主要作用是避免冲突。

go 标准包在 goroot 路径的 pkg 下，第三方包在 gopath 的 pkg 下，执行 go get 获取第三方包将把包安装在 gopath 中。

go package 中有 init 函数，用于在包导入时执行，如果仅仅只想执行 package 的 init 函数，可以 `import _ "fmt"`。

在 package 所在的文件目录执行 go install 将安装 package。

## pointer 

pointer 是指向某种类型的第一个地址，表示 memory 中的一个 address。

声明 var ptr *type.

赋值 *ptr = value。

new(T) 分配了一个零值填充的 T 类型的空间，并且返回其地址，一个 *T 类型的值，它返回的是一个指针。

make(T, args) 返回的是一个 T 类型。

type 用来创建自定义类型。

## embed type go

go 通过 embed type 来完成属性和方法的继承，go 中没有显示的继承，把一个类型嵌入另一个类型作为 field，而且可以是匿名 field 可以达到继承另一个 type 属性的目的。

被嵌入的 type 属性不能直接通过字面量进行初始化，必须是被嵌入的类型才可以。

```golang
t := Teacher{Person{name: "teacher"}}
```

## struct

go 的方法传值默认是按值传递，不能被修改，要想被修改必须要按照引用传递，所有 struct 的方法必须是 *p。

声明 struct 对象有 new、字面量、声明等不同的方法。

var c Circle, c := new(Circle) c = Circle{x:1,y:1}。

只有 new 返回的是 struct 指针，其他两种声明方式都不是。

## interface

go 的 interface 定义了一些列行为来达到多态。
通过实现 interface 的方法可以让不同的类型抽象成一个共同的基类，一个 struct 拥有 interface 的方法表示 struct 实现了 interface。
实现 interface 的方法 receiver 必须是 pointer， 不能是 value，因而 make 之后的值可以直接调用，字面量的值必须是 &value

```golang
//interface 定义了行为
type People interface {
    sayHello()
}

type Person struct {
    name string
}

func (p *Person) sayHello() {
    fmt.Println(p.name + " hello")
}

type Student struct {
    Person //匿名 field
}

type Teacher struct {
    Person
}

t := new(Teacher) 
t1 := Teacher{Person{name: "teacher1"}} //嵌入类型的属性必须通过嵌入类型进行初始化
```
t 实现了 people 的方法 sayHello，t1 并没有


## channel

channel 用来在 goroutine 间通信。

当尝试向 channel 发送数据，channel 必须等待 channel 的接收端是准备好的状态，channel 是用阻塞完成的。

channel 是有方向的 c chan<- string, c <-chan string。


## select 

select 类似于 switch case，select 会选择当前第一个 ready 状态的 channel，如果有多个同时准备好，会随机选一个，如果没有则阻塞直到有一个准备好为之。

select 可以有默认的 channel，如果在一定时间内没有任何 channel，select 将执行默认 case。