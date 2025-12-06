package main

import (
	"fmt"
	"net/http"
)

/*

a handler is an interface that has a method named ServeHTTP
with two parameters: an HTTPResponseWriter interface and
a pointer to a Request struct.

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

/*
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}
*/

//ServeMux has a method named ServeHTTP with same signature
//DefaultServeMux is an instance of ServeMux,and it is also
//an instance of the Handler struct

type MyHandler struct{}

// the ServeHTTP method does all the processing,so we need multiplexer
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
