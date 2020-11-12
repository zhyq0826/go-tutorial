package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

type Go struct {
	ID     int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
	GID    int    `gorm:"column:gid;type:int(11);"`
	M      int    `gorm:"column:m;type:int(11);"`
	Text   string `gorm:"column:text;"`
	Status string `gorm:"column:status;"`
}

func (Go) TableName() string {
	return "go"
}

func main() {
	initDB()
	f, err := os.Open("goroutine.txt")
	reader := bufio.NewReader(f)

	var s []byte
	var g Go
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		if string(buf) == "\n" {
			g.ID = 0
			g.Text = string(s)
			DB.Create(&g)
			s = []byte{}
			continue
		}
		line := string(buf)
		if strings.HasPrefix(line, "goroutine") {
			ss := strings.Split(line, " ")
			g.GID, _ = strconv.Atoi(ss[1])
			g.Status = strings.Join(ss[2:], " ")
			if strings.Contains(line, "minutes") {
				g.M, _ = strconv.Atoi(ss[len(ss)-2])
			}
		}

		s = append(s, buf...)
	}

	fmt.Println(err)
}
