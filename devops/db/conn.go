package db

import (
	"log"

	"github.com/jinzhu/gorm"

	// init mysql init
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB for database
var DB *gorm.DB

func init() {
	DB, err := gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// check db conn
	DB.Raw("select 1")
}
