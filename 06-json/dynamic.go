package main

import (
	"encoding/json"
	"fmt"
)

var jsonStrArray = `[
	{"id": 17, "username": "den", "phone": 0},
	{"id": "17", "address": "none", "company": "IT-company"}
]`

func main() {
	data := []byte(jsonStrArray)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("unpacked in empty interface:\n%#v\n\n", user1)

	user2 := map[string]interface{}{
		"id": 42,
		"username": "den",
	}
	var user2i interface{} = user2
	result, err := json.Marshal(user2i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string from map:\n %s\n", string(result))
}