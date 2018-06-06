package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// "github.com/zhyq0826/go-tutorial/devops/db"
	// "github.com/zhyq0826/go-tutorial/devops/db"

	// init mysql init
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB for database
var DB *gorm.DB

type S struct {
	ID int
}

func initDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(DB)

	// check db conn
	s := S{}
	DB.Raw("select id from domain").Scan(&s)
	log.Println(s)

}

// BaseModel for all devops model
type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
	ID        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
}

// Domain model
type Domain struct {
	BaseModel
	Name    string `gorm:"column:name;type:varchar(256)"`
	URL     string `gorm:"column:url;type:varchar(256)"`
	Private uint   `gorm:"column:private;type:tinyint;not null;default:0"`
}

// TableName
func (Domain) TableName() string {
	return "domain"
}

// QueryDomain domain list
func QueryDomain(page, limit int, name, url string) []Domain {
	if limit == 0 {
		limit = 10
	}
	domains := make([]Domain, limit)
	// var domains []Domain
	conditon := make([]interface{}, 0)
	query := ""
	if name != "" {
		query += "name = ?"
		conditon = append(conditon, name)
	}
	if url != "" {
		query += "and url = ?"
		conditon = append(conditon, url)
	}
	log.Println(DB)
	log.Println(query)
	log.Println(conditon)
	DB.Where(query, conditon...).Find(&domains)

	log.Println(domains)
	return domains
}

func main() {
	initDB()
	QueryDomain(0, 0, "name", "")
}
