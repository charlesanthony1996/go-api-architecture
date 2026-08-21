package main

import (
	"fmt"
	"net/http"
)

func ArticleHandler(w httpResponse.Writer, r *http.Request) {
	fmt.Fprintln(w, "Chained route works")
}

func main() {

}
