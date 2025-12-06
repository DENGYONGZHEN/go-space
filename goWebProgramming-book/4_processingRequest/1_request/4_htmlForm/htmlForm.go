package main

import (
	"fmt"
	"net/http"
)

func htmlForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)     //map[email:[deng@mail.com] firstName:[deng good] weather:[cloudy]]
	fmt.Fprintln(w, r.PostForm) //map[email:[deng@mail.com] firstName:[deng]]

}
func htmlFormMultiPart(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form) //map[firstName:[good] weather:[cloudy]]
	//PostForm only supports application/x-www-form-urlencoded
	fmt.Fprintln(w, r.PostForm) //	map[]

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/htmlForm", htmlForm)
	http.HandleFunc("/htmlFormMultiPart", htmlFormMultiPart)
	server.ListenAndServe()
}
