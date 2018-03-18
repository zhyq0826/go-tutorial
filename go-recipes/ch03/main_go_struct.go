package main

import (
	"fmt"
)

type Customer struct {
	Name, Email, Phone string
	Addresses          []Address
}

type Address struct {
	City string
}

func main() {
	var c Customer
	c.Name = "Jone"
	c.Email = "Jone@gmail.com"
	c.Phone = "010-10239877"
	fmt.Println("%#v", c)
	c2 := Customer{Name: "Jone", Email: "Jone@gmail.com", Phone: "010-10239877"}
	fmt.Println("%#v", c2)
	c3 := Customer{
		Name:  "Jone",
		Email: "Jone@gmail.com",
		Phone: "010-10239877",
		Addresses: []Address{
			Address{
				City: "San Francisco",
			},
		},
	}

	fmt.Println("%#v", c3)
}
