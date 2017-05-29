package main

import (
	"fmt"
)

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

type Speaker struct {
	Person
}

func (s *Speaker) sayHello() {
	fmt.Println("I am speaker")
}

func main() {
	s := new(Student)
	s.Person.name = "Jone"     //显示通过 person 初始化
	fmt.Println(s.name)        //隐式获取继承类型的属性
	fmt.Println(s.Person.name) //显示获取继承类型的属性
	s.sayHello()
	s.name = "student" //隐式
	fmt.Println(s.name)
	fmt.Println(s.Person.name)

	t := new(Teacher)
	t.name = "teacher"
	t.sayHello()

	//通过字面量声明的对象，并没有获得一个真正的指针
	t1 := Teacher{Person{name: "teacher1"}} //嵌入类型的属性必须通过嵌入类型进行初始化
	//Teacher does not implement People (sayHello method has pointer receiver)
	// append(class, t1)
	t1.sayHello()

	sp := Speaker{Person{name: "specker"}}
	sp.sayHello()

	//注意 sp 是一个 &sp，表示 sp 的一个指针
	slice := append(make([]People, 5), t, &sp) //append(class, t) evaluated but not used append 的放好值必须要接住
	//slice 长度是 5，只有一个非空值，其他是 nil 的 people
	for _, p := range slice {
		if p != nil {
			p.sayHello()
		}
	}
}
