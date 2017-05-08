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

- byte
- rune
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

go 中的字符串不同于 c，不是字符的序列。

go 中有单独的字符序列是 rune 数组。

多行字符串注意 + 号的位置。


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

append() 可以向 slice 追加元素并且自动扩容，扩容之后会返回新的底层 array，所以调用