package main

import (
	"fmt"
)

type People interface {
	SayHello()
	GetDetails()
}

type Person struct {
	name        string
	age         int
	city, phone string
}

//A person method
func (p Person) SayHello() {
	fmt.Printf("Hi, I am %s, from %s\n", p.name, p.city)
}

//A person method
func (p Person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", p.name, p.age, p.city, p.phone)
}

type Speaker struct {
	Person     //type embedding for composition
	speaksOn   []string
	pastEvents []string
}

type Organizer struct {
	Person  //type embedding for composition
	meetups []string
}

type Attendee struct {
	Person    //type embedding for composition
	interests []string
}

func main() {

}
