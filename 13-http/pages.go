package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Single page:", r.URL.String())
	})

	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Multiple page:", r.URL.String())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Main page")
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
