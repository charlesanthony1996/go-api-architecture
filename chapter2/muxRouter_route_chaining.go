package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Chained route works")
}

func main() {
	r := mux.NewRouter()
	r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(ArticleHandler)
	srv := &http.Server{
		Handler: r, Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second, ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
