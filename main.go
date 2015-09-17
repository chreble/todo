package main

import (
	"github.com/chreble/todo/server"
	"log"
	"net/http"
)

func main() {
	rt := server.RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8080", rt))
}
