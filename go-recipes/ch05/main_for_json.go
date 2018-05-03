package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID                            int
	FirstName, LastName, JobTitle string
}

func main() {
	emp := Employee{
		ID:        100,
		FirstName: "Shiju",
		LastName:  "Varghese",
		JobTitle:  "Architect",
	}
	// Encoding to JSON
	data, _ := json.Marshal(emp)
	fmt.Println(string(data))

	b := []byte(`{"ID":101,"FirstName":"Irene","LastName":"Rose","JobTitle":"Developer"}`)

	emp1 := Employee{}
	json.Unmarshal(b, &emp1)
	fmt.Println(emp1)
}
