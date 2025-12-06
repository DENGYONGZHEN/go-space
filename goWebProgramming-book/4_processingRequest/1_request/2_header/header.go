package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	//[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7]
	h := r.Header["Accept"]
	fmt.Fprintln(w, h)
	//gzip, deflate, br, zstd
	e := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, e)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/header", headers)
	server.ListenAndServe()
}
