package main

import (
	"io"
	"log"
	"net/http"
)

func myServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", myServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
