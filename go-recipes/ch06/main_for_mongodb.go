package main

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	session, _ := mgo.Dial("localhost")
	fmt.Println(session.DB("test").C("test"))
}
