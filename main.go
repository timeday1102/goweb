package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/timeday1102/goweb/handle"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
	w.Write([]byte("hello"))
}

func main() {

	connStr := "user=test password=test54Pt8@52#1 dbname=postgres port=19845 host=101.43.34.22 sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected!")
	}

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
