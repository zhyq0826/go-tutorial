package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item : %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	// http.HandlerFunc is type this call http.HandlerFunc(db.list) is a conversion
	// http.handlerFunc 是一个 函数 type ，有 serveHTTP 方法 满足 http.Handler interface
	// type HandlerFunc func(w ResponseWriter, r *Request)
	// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	//         f(w, r)
	// }
	// 转换之后就是满足是一个 handler
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
