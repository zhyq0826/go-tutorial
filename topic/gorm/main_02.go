package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"fmt"

	// init mysql init
	"time"

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

func query() (ret Domain){
	DB.Where("id = ?", 1).First(&ret)
	return
}

func query2() (ret Domain){
	DB.Where("id = ?", 1).First(&ret)
	return
}

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
	// var cond map[string]interface{}
	// DB.Model(Domain{}).Where(cond).First(&Domain{})
	var ds []Domain
	DB.Exec("select * from domain").Where("id=?", 1).Order("id desc").Find(&ds) //SELECT * FROM `domains`  WHERE (id=1) ORDER BY id desc
	DB.Delete(&Domain{ID: 0})                                                   //DELETE FROM `domains`
	DB.Where(map[string]interface{}{
		"name": "",
		"url":  "",
	}).Delete(Domain{}) // DELETE FROM `domains`  WHERE (`domains`.`name` = '') AND (`domains`.`url` = '')

	DB.Where("name = ? and url = ?", "", "").Delete(Domain{}) //DELETE FROM `domains`  WHERE (name = '' and url = '')

	DB.Model(&Domain{}).Where(map[string]interface{}{
		"name": "",
		"url":  "",
	}).First(&Domain{}) //SELECT * FROM `domains`  WHERE (`domains`.`name` = '') AND (`domains`.`url` = '') ORDER BY `domains`.`id` ASC LIMIT 1

	DB.Where("id = 1").Delete(&Domain{ID: 0})                     //DELETE FROM `domains`  WHERE (id = 1)
	DB.Where("id = 1").Delete(&Domain{ID: 1})                     //ELETE FROM `domains`  WHERE `domains`.`id` = 1 AND ((id = 1))
	DB.Where(1).Delete(Domain{})                                  //DELETE FROM `domains`  WHERE (`domains`.`id` = 1)
	DB.Where(0).Delete(Domain{})                                  //DELETE FROM `domains`  WHERE (`domains`.`id` = 0)
	DB.Model(Domain{}).Where(0).Update("name", "name")            //UPDATE `domains` SET `name` = 'name', `updated_at` = '2019-09-26 17:22:08'  WHERE (`domains`.`id` = 0)
	DB.Where(0).Delete(Domain{ID: 1})                             //DELETE FROM `domains`  WHERE `domains`.`id` = 1 AND ((`domains`.`id` = 0))
	DB.Model(Domain{ID: 1}).Where(2).Update("name", "name")       // UPDATE `domains` SET `name` = 'name', `updated_at` = '2019-09-26 17:36:44'  WHERE `domains`.`id` = 1 AND ((`domains`.`id` = 2))
	DB.Model(Domain{ID: 0}).Where(2).Update(Domain{Name: "name"}) // UPDATE `domains` SET `name` = 'name', `updated_at` = '2019-09-26 17:36:44'  WHERE `domains`.`id` = 1 AND ((`domains`.`id` = 2))
	DB.Table("domain").Where(2).Update(Domain{Name: "name"})      // UPDATE `domain` SET `name` = 'name'  WHERE (`domain`.`` = 2)
	DB.Model(Domain{ID: 0}).Where(2).Update(Domain{Name: "name"}) // UPDATE `domains` SET `id` = 2, `name` = 'name', `updated_at` = '2019-09-26 18:31:12'  WHERE (`domains`.`id` = 2)

	DB.Model(Domain{}).Where(0).Update("name", "name") //UPDATE `domains` SET `name` = 'name', `updated_at` = '2019-09-26 18:38:00'  WHERE (`domains`.`id` = 0)
	DB.Where(0).Delete(Domain{})                       //DELETE FROM `domains`  WHERE (`domains`.`id` = 0)
	DB.Model(Domain{}).Where("c1 = ? and c2 = ?", 1, 2).Update("name", "name")
	DB.Where("c1 = ? and c2 = ?", 1, 2).Delete(Domain{})  //DELETE FROM `domains`  WHERE (c1 = 1 and c2 = 2 )
	DB.Table("domains").Select("c1, c2").Where(1).Rows()  //SELECT c1, c2 FROM `domains`  WHERE (`domains`.`` = 1)
	DB.Model(Domain{}).Select("c1, c2").Where(1).Rows()   //SELECT c1, c2 FROM `domains`  WHERE (`domains`.`` = 1)
	DB.Where("c1 = ? and c2 = ?", 1, 2).First(&Domain{})  //SELECT * FROM `domains`  WHERE (c1 = 1 and c2 = 2) ORDER BY `domains`.`id` ASC LIMIT 1
	DB.Where(0).First(&Domain{})                          //SELECT * FROM `domains`  WHERE (`domains`.`id` = 0) ORDER BY `domains`.`id` ASC LIMIT 1
	DB.Where("c1 = ? and c2 = ?", 1, 2).Find(&[]Domain{}) //SELECT * FROM `domains`  WHERE (c1 = 1 and c2 = 2)
	fmt.Println(query())
}
