package main

import (
	"encoding/json"
	"fmt"
)

//Models for JSON resources
type (
	// UserResource For Post - /user/register
	UserResource struct {
		Data UserModel `json:"data"`
	}

	// UserModel reperesents a user
	UserModel struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
)

func main() {
	body := []byte(`{"data": {"lastname": "san", "password": "1123456", "email": "zhyq@gmail.com", "firstname": "zhyq"}}`)

	var dataResource UserResource
	json.Unmarshal(body, &dataResource)
	fmt.Println(dataResource)
}
