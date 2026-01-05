package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.writeHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id: %s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", query)
}
