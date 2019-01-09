package main

import (
	"encoding/json"
	"fmt"
)

type UserProfile struct {
	ID 			int `json:"user_id,string"`
	Username 	string
	Address 	string `json:",omitempty"`
	Company		string `json:"-"`
}

func main() {
	up := &UserProfile{
		ID: 42,
		Username: "den",
		Address: "",
		Company: "IT-company",
	}

	result, err := json.Marshal(up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string:\n\t%s\n", string(result))
}