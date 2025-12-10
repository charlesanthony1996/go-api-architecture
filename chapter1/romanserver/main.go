package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/charles/go-api-arch/romanNumerals"
)

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// if request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))

			if number == 0 || number > 10 {
				// if response is not in the list, send not found status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - not found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - bad request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
