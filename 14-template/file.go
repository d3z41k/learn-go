package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))

	users := []User{
		User{1, "Denis", true},
		User{2, "<i>Ivan</i>", false},
		User{3, "Vasily", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tmpl.Execute(w,
			struct {
				Users []User
			}{
				users,
			})
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
