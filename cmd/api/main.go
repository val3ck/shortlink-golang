package main

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is my content.")
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Body)
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8080", nil)
}
