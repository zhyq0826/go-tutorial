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