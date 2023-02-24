package main

import (
	"fmt"
	"net/http"
)

type hellohandler struct{}

func (mh *hellohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type abouthandler struct{}

func (mh *abouthandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func parse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	for key, val := range r.Form {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("val: %v\n", val)
	}
	fmt.Println(r.Header)

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	fmt.Println(string(body))

	fmt.Fprintf(w, "parse successful")
}

func main() {

	// http.ListenAndServe("localhost:8080", nil) // DefaultServeMux

	helloHd := &hellohandler{}
	aboutHd := &abouthandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // DefaultServeMux
	}
	http.Handle("/hello", helloHd)
	http.Handle("/about", aboutHd)

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	// 下面等价
	// http.HandleFunc("/welcome", welcome)
	http.Handle("/welcome", http.HandlerFunc(welcome))
	http.HandleFunc("/parse", parse)
	http.Handle("/", http.FileServer(http.Dir("wwwroot")))

	server.ListenAndServe()
}
