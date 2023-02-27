package main

import (
	"net/http"

	"github.com/timeday1102/goweb/handle"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
	w.Write([]byte("hello"))
}

func main() {

	// http.ListenAndServe("localhost:8080", nil) // DefaultServeMux

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // DefaultServeMux
	}
	// 下面等价
	// http.HandleFunc("/welcome", welcome)
	http.Handle("/welcome", http.HandlerFunc(welcome))
	http.HandleFunc("/parse", handle.Parse)
	http.Handle("/", http.FileServer(http.Dir("wwwroot")))
	http.HandleFunc("/login", handle.Login)
	http.HandleFunc("/template", handle.Template)
	server.ListenAndServe()
}
