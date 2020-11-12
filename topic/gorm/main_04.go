package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB for database
var DB *gorm.DB

func initDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@/blog?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(DB)
}

func main() {
	initDB()
	rows, _ := DB.Raw("show tables").Rows()
	defer rows.Close()
	var name = ""
	for rows.Next() {
		rows.Scan(&name)
		fmt.Println(name)
	}

	DB.Table(name).Delete(nil, "id in (?)", []int{0})

}
