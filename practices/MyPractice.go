package main

import "fmt"

type User struct {
	id   int
	name string
	pw   string
}

func (u User) String() string { //ToString метод в джаве
	return fmt.Sprintf("name is %s", u.name)
}

func main() {
	u := User{
		id:   1,
		name: "Babo",
		pw:   "123123",
	}

	fmt.Println(u)
}
