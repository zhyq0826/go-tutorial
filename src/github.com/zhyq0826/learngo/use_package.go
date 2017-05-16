package main

import "time" //单行导入
import (
	"fmt" //分组
)

// import (
//  "gopkg.in/mgo.v2"
// )

import (
	m "math" //别名
)

import (
	"github.com/zhyq0826/mypack"
)

func main() {
	fmt.Println(10)
	time.Now()
	fmt.Println(m.Pi)
	// session, err := mgo.Dial("mongodb://localhost:27001")
	// c := session.DB("test").C("test")
	// err := c.Find("").One(&result)
	// fmt.Println(result)
	fmt.Println(mypack.Even(10))
}
