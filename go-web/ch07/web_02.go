package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"path"
	"strconv"
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

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}
