package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello !")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world !")
}

// HandleFunc registers the handler function for the given pattern in [DefaultServeMux].
func main() {
	server := http.Server{
		Addr: "12.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}
