# type

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


# variable 

go 是类型置后的。

可选的 `;` 一般来说如果一行只有一条语句 `;` 是不需要的，如果要在一行出现多条语句 `;` 是必须的。

go 可以动态进行类型推断，如果有初始化值可以不必一开始就声明变量类型。

`:=` 赋值只能用在 function 中，表示声明并初始化。

**declare variable**

var variable_list optional_data_type


# constant

常量在 go 中只能是 boolean、string、numerical。

常量可以使用枚举、指定初始值、或者显式指定类型和初始值。


# operator

go 中只能同类型的进行操作，不能混用，也就是不能用 int32 + int16


两个重要的操作符是 &、* ，同 c。

```
var ptr *int  //声明指针
var a = 0

ptr = &a  //指针赋值
```

# string

go 中字符串是不可变的。

go 中的字符串不同于 c，不是字符的序列，而是一个个 byte，直接输出 s[0] 是 byte 代表的 10 进制数

go 中有单独的字符序列是 rune 数组。

多行字符串注意 + 号的位置。

# scope

go 是 block scope language


# slice 

slice 是 array 的一段，唯一不同的是可变。

创建 slice 可以用 make，array[start:end], []type 等方式

不指定长度就是直接创建 slice: var sl []int。

append 的返回值必须要接住，否则报：append(class, t) evaluated but not used。

make([]int, 5) 表示创建长度为 5 被 0 值填充的一个 slice。


# if switch

go 的控制结构少，有 if、for switch。

if、for 接受初始化语句。

switch 不同于其他语言，不需要 break，因为匹配失败之后不会继续向下执行，可以用 fallthrough。

if 的大括号必须在同一行。


# goto for

go 有 goto 语句，标签名是大小写敏感的。

go 没有 while 语句，可以使用 for 替代。

for 有 break 和 continue 语句，break 可以指定标签来跳出哪个循环

# range 

range 对 array、slice、map 的用法

# array

声明 array 必须要有个数说明。

slice 是一个指向底层 array 的类型，即使从新的 slice 创建 slice，两个 slice 也是指向同一个 array

append() 可以向 slice 追加元素并且自动扩容，扩容之后会返回新的底层 array，所以调用 append 可能会产生新的底层 array。


# function

函数的定义
```
func (p mytype) functioname(q int) (r,s int) {return 0, 0}
```

socpe 遵循局部覆盖全部的原则。

如果函数有返回值，那 return 语句是必须的。

函数可接受变参，参数类型一致的参数列表，其实是一个 slice。

函数可以作为值来使用，可以把函数赋值给一个变量。

# defer

defer 在 go 中可以让函数延迟执行，主要用来做一些清理工作，比如 socket 的关闭，文件资源的释放，数据库的关闭等。

defer 执行的必须是函数的执行，匿名函数屁股后面一定是调用符号()。

defer 表示在包裹函数执行完成之后在执行，所以 defer 可以用来恢复异常，即 recover 在 defer 函数中使用才有意义。

多个 defer 函数 是按照后进先出的顺序执行，是一个栈操作。

函数 defer 时，被 defer 的函数调用的参数首先被 evaluated，也就是 defer 函数的参数是在 defer 时就执行固定的，即使这些参数在 defer 之后更改也没有影响。



# package

go package 是 go 源文件的 package 名称，和 源文件所在 directory 其实是没有关系的，不过一般意义上 path 的最后一个元素和源文件的 package 
名称相同。

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand


像 Python 一样，导入可单行，也可分组。

package 默认只导出首字母大写的变量、函数，小写不导出，也就是可以使用那些头字母是大写的函数。

包可以有别名，别名的主要作用是避免冲突。

go 标准包在 goroot 路径的 pkg 下，第三方包在 gopath 的 pkg 下，执行 go get 获取第三方包将把包安装在 gopath 中。

go package 中有 init 函数，用于在包导入时执行，如果仅仅只想执行 package 的 init 函数，可以 `import _ "fmt"`。

在 package 所在的文件目录执行 go install 将安装 package。

# pointer 

pointer 是指向某种类型的第一个地址，表示 memory 中的一个 address。

```
声明 var ptr *type.

赋值 *ptr = value。

new(T) 分配了一个零值填充的 T 类型的空间，并且返回其地址，一个 *T 类型的值，它返回的是一个指针。
```

make(T, args) 返回的是一个 T 类型。

type 用来创建自定义类型。

# embed type go

go 通过 embed type 来完成属性和方法的继承，go 中没有显示的继承，把一个类型嵌入另一个类型作为 field，而且可以是匿名 field 可以达到继承另一个 type 属性的目的。

被嵌入的 type 属性不能直接通过字面量进行初始化，必须是被嵌入的类型才可以。

```golang
t := Teacher{Person{name: "teacher"}}
```

# struct

go 的方法传值默认是按值传递，不能被修改，要想被修改必须要按照引用传递，所有 struct 的方法必须是 `*p`。

声明 struct 对象有 new、字面量、声明等不同的方法。

var c Circle, c := new(Circle) c = Circle{x:1,y:1}。

只有 new 返回的是 struct 指针，其他两种声明方式都不是。

# interface

go 的 interface 定义了一些列行为来达到多态。

interface类型定义了一组方法，如果某个 struct 实现了某个接口的所有方法，则此对象就实现了此接口。


任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。


空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的`void*`类型。


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


This is not pure duck typing, because when possible the Go compiler will statically check whether the type implements the interface. However, Go does have a purely dynamic aspect, in that you can convert from one interface type to another. In the general case, that conversion is checked at run time. If the conversion is invalid -- if the type of the value stored in the existing interface value does not satisfy the interface to which it is being converted -- the program will fail with a run time error.

go 编译器会在编译时静态检查是否一个类型实现了某个接口，运行时 go 又可以帮助你把一种类型的 interface 转换成另一个中类型，比如某个类型实现了多个 interface，在不同的调用场景下使用不同类型的 interface 的执行，这种 interface 之间的转换，go 已经帮你做了，一旦转换发生错误，go 发现转换无法完成，就会抛运行时错误。


The receiver type must be of the form T or `*T` where T is a type name. T is called the receiver base type or just base type. The base type must not be a pointer or interface type and must be declared in the same package as the method.

receiver type 只能是 T 类型或者 T 类型的指针 `*T`。T 被称之为 base type，base type 不能是指针或者是 interface type，必须在当前包中声明。


interface 类型是一组方法的集合但不包含方法的实现，也就是 interface 仅仅只是一个定义，因而 interface 不能作为 receiver


## 推断 interface 中的值


Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T) 这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。


```go
if value, ok := element.(int); ok {
    
}

```

switch element.(type) 判断


```

switch value := element.(type){
    case int:
        fmt.Println("")
    default:
        fmt.Println("")
}

```

## interface value 进行比较

两个 interface value 值进行比较，必须他们的 dynamic type 是相同的，而且 dynamic value 也是相同的才是相同的，前提条件必须是 dynamic value 是可比较的，如果是不可以比较的就会 panic，所以在比较 interface 的 value 或者 switch 语句中使用 interface  一定要注意 拍你才

# channel

channel 用来在 goroutine 间通信。

当尝试向 channel 发送数据，channel 必须等待 channel 的接收端是准备好的状态，channel 是用阻塞完成的。

channel 是有方向的 c chan<- string, c <-chan string。

If the channel is unbuffered, the sender blocks until the receiver has received the value. If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value


# select 

select 类似于 switch case，select 会选择当前第一个 ready 状态的 channel，如果有多个同时准备好，会随机选一个，如果没有则阻塞直到有一个准备好为之。

select 可以有默认的 channel，如果在一定时间内没有任何 channel，select 将执行默认 case。

# make 和 new 的区别

new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type `*T.`

## new(T) 返回的是 T 类的指针

new(T) 为一个 T 类型新值分配空间并初始化为 T 的零值，返回新值的地址，也就是 T 类型的指针 `*T`。

```go
p1 := new(int)
fmt.Printf("p1 --> %#v \n ", p1) //(*int)(0xc42000e250) 
fmt.Printf("p1 point to --> %#v \n ", *p1) //0

var p2 *int
i := 0
p2 = &i
fmt.Printf("p2 --> %#v \n ", p2) //(*int)(0xc42000e278) 
fmt.Printf("p2 point to --> %#v \n ", *p2) //0 

```

上面的代码是等价的，new(int) 初始化为 int 的零值，也就是 0，并返回 T 的指针，这和直接声明指针并初始化的效果是相同的。

## make 只能用于 slice,map,channel

make 只能用于 slice，map，channel 三种类型，make(T, args) 返回的是初始化之后的 T 类型的值，这个新值并不是 T 类型的零值，也不是指针 `*T`，是 T 的引用。

```go
var s1 []int
if s1 == nil {
    fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
}

s2 := make([]int, 3)
if s2 == nil {
    fmt.Printf("s2 is nil --> %#v \n ", s2)
} else {
    fmt.Printf("s2 is not nill --> %#v \n ", s2)// []int{0, 0, 0}
}

```
slice 的零值是 nil，使用 make 之后 slice 是一个初始化的 slice，即 slice 的内容被类型 int 的零值填充，形式是 [0 0 0]

map 和 channel 也是类似的。
```go
var m1 map[int]string
if m1 == nil {
    fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
}

m2 := make(map[int]string)
if m2 == nil {
    fmt.Printf("m2 is nil --> %#v \n ", m2)
} else {
    fmt.Printf("m2 is not nill --> %#v \n ", m2) map[int]string{} 
}


var c1 chan string
if c1 == nil {
    fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
}

c2 := make(chan string)
if c2 == nil {
    fmt.Printf("c2 is nil --> %#v \n ", c2)
} else {
    fmt.Printf("c2 is not nill --> %#v \n ", c2)//(chan string)(0xc420016120)
}
```

## make(T, args) 返回的是 T 的 引用

如果不特殊声明，go 的函数默认都是按值穿参，即通过函数传递的参数是值的副本，在函数内部对值修改不影响值的本身，但是 make(T, args) 返回的值通过函数传递参数之后可以直接修改。

```go
func modifySlice(s []int) {
    s[0] = 1
}

s2 := make([]int, 3)
fmt.Printf("%#v", s2) //[]int{0, 0, 0}
modifySlice(s2)
fmt.Printf("%#v", s2) //[]int{1, 0, 0}

```

这说明 make(T, args) 返回的是引用类型，在函数内部可以直接更改原始值，对 map 和 channel 也是如此。

```go
func modifyMap(m map[int]string) {
    m[0] = "string"
}

func modifyChan(c chan string) {
    c <- "string"
}

m2 := make(map[int]string)
if m2 == nil {
    fmt.Printf("m2 is nil --> %#v \n ", m2) 
} else {
    fmt.Printf("m2 is not nill --> %#v \n ", m2) //map[int]string{}
}

modifyMap(m2)
fmt.Printf("m2 is not nill --> %#v \n ", m2) // map[int]string{0:"string"}


c2 := make(chan string)
if c2 == nil {
    fmt.Printf("c2 is nil --> %#v \n ", c2)
} else {
    fmt.Printf("c2 is not nill --> %#v \n ", c2)
}

go modifyChan(c2)
fmt.Printf("c2 is not nill --> %#v ", <-c2) //"string"

```

## go 中的面向对象

- https://thenewstack.io/understanding-golang-type-system/

定义 Person 作为我们的基础属性集合 struct，定义 People 表示 Person 的行为，
Person 实现了 People 的方法也就实现了 People 接口。

Speaker，Organizers，Attendee 通过嵌入类型继承了 Person 的属性和方法，也就是实现 People
接口，他们都可以用 People 类型来表示，通过 Meetup 中的 People slice 来
表示，Organizers 和 Speaker 都复写了 Getdetails 方法。
