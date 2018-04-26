package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@/blog")
	if err != nil {
		panic(err)
	}
}

type Post struct {
	Id      int    `json: "id"`
	Title   string `json: "title"`
	Uid     int    `json: "uid"`
	Content string `json: "content"`
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	statment, err := db.Prepare("select id, title, uid, text from entries where id = ?")
	defer statment.Close()
	statment.QueryRow(id).Scan(&post.Id, &post.Title, &post.Uid, &post.Content)
	return
}

func (p *Post) Create() (err error) {
	statement, _ := db.Prepare("INSERT INTO entries (id, title, uid, dig_count, like_count, comment_count, text, atime) VALUES (0, ?, ?, ?, ?, ?, ?, ?)")
	defer statement.Close()
	// 等号左边至少有一个是未声明的变量
	res, err := statement.Exec(p.Title, p.Uid, 0, 0, 0, "text", time.Now())
	id, _ := res.LastInsertId()
	p.Id = int(id)
	return
}

func (p *Post) Update() (err error) {
	statement, _ := db.Prepare("update entries set title = ? where id = ?")
	defer statement.Close()
	statement.Exec(p.Title, p.Id)
	return
}

func (p *Post) Delete() (err error) {
	statement, err := db.Prepare("delete from entries where id = ?")
	statement.Exec(p.Id)
	return
}
