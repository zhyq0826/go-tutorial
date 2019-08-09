package main

import (
	"log"

	"github.com/jinzhu/gorm"

	// init mysql init
	"time"

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

// Domain model
type Domain struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
	ID        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
	Name      string    `gorm:"column:name;type:varchar(256)"`
	URL       string    `gorm:"column:url;type:varchar(256)"`
	Private   uint      `gorm:"column:private;type:tinyint;not null;default:0"`
}

// TableName
// func (Domain) TableName() string {
// 	return "domain"
// }

func main() {
	initDB()

	DB.First(&Domain{})                 // SELECT * FROM `domain`   ORDER BY `domain`.`id` ASC LIMIT 1
	DB.Delete(&Domain{}, "name = ?", 1) // DELETE FROM `domain`  WHERE (name = 1)
	// Model 调用形式只认主键
	DB.Model(&Domain{ID: 1, Name: "name"}).Update("name", "name") // UPDATE `domain` SET `name` = 'name', `updated_at` = '2019-07-26 18:01:57'  WHERE `domain`.`id` = 1
	// 如果没有主键，更新的是所有记录,要尽可能屏蔽这个操作
	DB.Model(&Domain{}).Update("name", "name")                         // UPDATE `domain` SET `name` = 'name', `updated_at` = '2019-07-26 18:20:57'
	DB.Model(&Domain{}).Where("id = ?", 1).Update("name", "name")      // UPDATE `domain` SET `name` = 'name', `updated_at` = '2019-07-26 18:20:57'  WHERE (id = 1)
	DB.Model(&Domain{ID: 1}).Where("id = ?", 1).Update("name", "name") // UPDATE `domain` SET `name` = 'name', `updated_at` = '2019-07-26 18:24:42'  WHERE `domain`.`id` = 1 AND ((id = 1))
	DB.Model(Domain{}).Where("id = ?", 1).Update("name", "name")       // UPDATE `domain` SET `name` = 'name', `updated_at` = '2019-07-26 19:00:08'  WHERE (id = 1)
}
