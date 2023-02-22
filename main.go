package main

import "net/http"

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
	server.ListenAndServe()

}
