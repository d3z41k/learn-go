package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID 			int
	Username 	string
	phone 		string // private
}

var jsonStr = `{"id": 42, "username": "den", "phone": "8-999-999-99-99"}`

func main() {
	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("strunct:\n\t%#v\n\n", u)

	u.phone = "8-999-999-99-90"
	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string:\n\t%s\n", string(result))
}
