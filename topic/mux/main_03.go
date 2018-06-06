package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// 如果 content-type 是 json 但是没有进行 json 序列化，直接 write 会变成下载文件，response 头部变成 oct-stream
	vars := mux.Vars(req)
	data := make([]byte, 10)
	for k, v := range vars {
		data = append(data, k...)
		data = append(data, "="...)
		data = append(data, v...)
		data = append(data, "&"...)
	}
	res.WriteHeader(http.StatusOK)
	// res.Header().Set("Content-Type", "text/html")
	res.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(string(data))
	// res.Write(data)
	res.Write(jsonData)
	// fmt.Fprintf(res, string(data))
}

func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/product").Subrouter()
	s.HandleFunc("/", indexHandler)
	http.Handle("/", r)
	r.Use(simpleMw)
	http.ListenAndServe("localhost:8300", nil)
}
