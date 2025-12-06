package main

import (
	"fmt"
	"net/http"
)

func url(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	fmt.Fprintln(w, url)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/url", url)
	server.ListenAndServe()
}
