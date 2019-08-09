package main

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB for database
var DB *gorm.DB

func initDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(DB)
}

func main() {
	initDB()
	DB.Exec("CREATE DATABASE `saber_test2` /*!40100 DEFAULT CHARACTER SET utf8mb4 */")
}
