package components

var _ User = &user{}

type user struct {
	Email, Username, Password string
	Age                       byte
}

func (u *user) ChangePW(pw string) {
	u.Password = pw
}

func (u *user) ChangeAge(age byte) {
	u.Age = age
}

type User interface {
	ChangePW(pw string)
	ChangeAge(age byte)
}

func NewUser(email, username, password string) User { //Реализация от интерфейса (по сути конструктор)
	u := user{
		Email:    email,
		Username: username,
		Password: password,
	}

	return &u
}
