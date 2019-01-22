package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("RequestID", "jdf83ndu2jb2ubd6vb5sr")

	fmt.Fprintln(w, "You browser is", r.UserAgent())
	fmt.Fprintln(w, "You accept", r.Header.Get("Accept"))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}