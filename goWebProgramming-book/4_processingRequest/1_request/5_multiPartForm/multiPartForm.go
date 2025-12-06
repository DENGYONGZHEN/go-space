package main

import (
	"fmt"
	"net/http"
)

func multiPartForm(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(1024)
	//fmt.Fprintln(w, r.MultipartForm) //&{map[email:[deng@mail.com] firstName:[deng]] map[]}
	fmt.Fprintln(w, r.FormValue("email")) //deng@mail.com
	fmt.Fprintln(w, r.Form)               //map[email:[deng@mail.com] firstName:[good deng] weather:[cloudy]]
	fmt.Fprintln(w, r.PostForm)           //map[email:[deng@mail.com] firstName:[deng]]
	fmt.Fprintln(w, r.MultipartForm)      //&{map[email:[deng@mail.com] firstName:[deng]] map[]}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/multiPartForm", multiPartForm)
	server.ListenAndServe()
}
