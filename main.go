package main

import (
    "net/http"
    "log"
    "github.com/chreble/todo/server"
)

func main() {
    rt := server.RegisterHandlers()
    log.Fatal(http.ListenAndServe(":8080", rt))
}