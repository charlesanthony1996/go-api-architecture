package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Settings for article %s\n", vars["id"])

}
