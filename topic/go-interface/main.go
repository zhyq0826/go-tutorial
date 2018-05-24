package main

import "fmt"

// Animal  for
type Animal interface {
	Speak() string
}

// Dog  for
type Dog struct {
}

// Speak  method
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat  for
type Cat struct {
}

// Speak  for
func (c Cat) Speak() string {
	return "Meow!"
}

// Llama   for
type Llama struct {
}

// Speak for
func (l Llama) Speak() string {
	return "?????"
}

//JavaProgrammer  for
type JavaProgrammer struct {
}

//Speak   for
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
