package main

import (
	"log"

	"github.com/travisjeffery/d-fuji/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
