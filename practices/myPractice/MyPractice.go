package myPractice

import "fmt"

type User struct {
	Id   int
	Name string
	Pw   string
	err  error
}

func (u *User) String() string { //ToString метод в джаве
	return fmt.Sprintf("user's name is \"%s\"", u.Name)
}

func (u *User) Error() string {
	return u.err.Error()
}
