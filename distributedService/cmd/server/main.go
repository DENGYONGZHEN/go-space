package main

import (
	"log"

	"github.com/deng/go-space/little-practices/distributedServices/internal/server"
)

func main() {
	server := server.NewHTTPServer(":8080")
	log.Fatal(server.ListenAndServe())
}
