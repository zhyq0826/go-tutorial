package main

import (
	"fmt"
)

type Customer struct {
	Name, Email, Phone string
	Addresses          []Address
}

func (c *Customer) ToAddress() string {
	// c.Addresses == (*c).Addresses
	for _, v := range c.Addresses {
		return fmt.Sprintf("city is %s", v.City)
	}
	return ""
}

func (c *Customer) ChangeEmail(email string) {
	c.Email = email
}

func (c Customer) ChangeName(name string) {
	c.Name = name
}

type Address struct {
	City string
}

func main() {
	c := &Customer{
		Name:  "Jone",
		Email: "Jone@gmail.com",
		Phone: "010-10239877",
		Addresses: []Address{
			{
				City: "San Francisco",
			},
		},
	}
	fmt.Println("before change email", c.Email)
	c.ChangeEmail("Pone@gmail.com")
	fmt.Println("after change email", c.Email)

	// 如果 receiver 不是 pointer go 会自动转换 (*c).ChangeName
	fmt.Println("before change name", c.Name)
	c.ChangeName("Pone")
	fmt.Println("after change name", c.Name)
}
