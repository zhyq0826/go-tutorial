package main

import (
	"fmt"
	"time"
)

type Employe struct {
	Id       int
	Name     string
	Salary   int
	DoB      time.Time
	Position string
	// invalid recursive type Employe
	// NewEmploye Employe
	NewEmploye *Employe
	// 聚合 type 无法包含自身，但是可以是自身的 pointer
}

var dilbert Employe

func EmployeId(id int) *Employe {
	return &dilbert
}

func main() {
	// dilbert is a variable so can assigne
	dilbert.Salary += 500
	position := &dilbert.Position
	*position = "Senior" + *position
	fmt.Println("Salary is ", dilbert.Salary, "Position is ", dilbert.Position)

	var employeOfTheMonth *Employe = &dilbert
	// 等价于 (*employeOfTheMonth).Position += "proactoive team leader"
	employeOfTheMonth.Position += "proactoive team leader"
	fmt.Println("Salary is ", dilbert.Salary, "Position is ", dilbert.Position)
	// 如果返回是 Employe cannot assign to EmployeId(1).Salary
	// 因为左边结果无法被 identity as a variable
	EmployeId(1).Salary = 0
}
