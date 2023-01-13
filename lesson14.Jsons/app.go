package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Age       byte   `json:"age"`
	Pw        string `json:"pw"`
	IsBlocked bool   `json:"is_blocked"`
	Roles     []Role `json:"roles"`
}

type Role struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

func main() {
	//serialize()

	const namePath = "resources/package.json"

	var dat User

	data, err := os.ReadFile(namePath)
	if err != nil {
		panic(err)
	}

	if err2 := json.Unmarshal(data, &dat); err2 != nil {
		panic(err2)
	}

	fmt.Println(dat)
}

func serialize() {
	var roles []Role
	role1 := Role{
		Id:   100,
		Role: "admin",
	}
	role2 := Role{
		Id:   101,
		Role: "user",
	}
	roles = append(roles, role1, role2)

	sv := User{
		Name:      "alex",
		Username:  "bobo",
		Age:       18,
		Pw:        "albobo2005",
		IsBlocked: false,
		Roles:     roles,
	}

	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))
}
