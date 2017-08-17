package main

import (
	"fmt"
)

// student 行为
type StudentTalk interface {
	talk()
}

// teacher 行为
type TeacherTalk interface {
	say()
}

// people 行为
type PeopleTalk interface {
	StudentTalk
	TeacherTalk
}

// Person 内嵌了 两个 interface，这两个 interface 定义了 PeopleTalk 的行为，也就是说如果 Person 实现了这两个行为， Person 也就实现了 PeopleTalk
type Person struct {
	TeacherTalk
	StudentTalk
}

type Student struct{}

func (s *Student) talk() {
	fmt.Println("student talk")
}

type Teacher struct{}

func (t *Teacher) say() {
	fmt.Println("teacher say")
}

//实现了 PeopleTalk 必然实现了 StudentTalk 和 TeacherTalk
func meet(p PeopleTalk) {
	fmt.Println("====>people meet<====")
	meetTeacher(p)
	meetStudent(p)
}

func meetTeacher(ps TeacherTalk) {
	fmt.Println("====>teacher meet<====")
	ps.say()
}

func meetStudent(ps StudentTalk) {
	fmt.Println("====>student meet<====")
	ps.talk()
}

func main() {
	t := Teacher{}
	s := Student{}
	// Person 被实
	p := Person{TeacherTalk: &t, StudentTalk: &s}
	meet(p)
}
