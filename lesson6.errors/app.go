package main

import "fmt"

//Errors

type customError struct {
	Err    error
	Custom string
	Field  int
}

func (e *customError) Error() string {
	return e.Err.Error()
}

func (e *customError) Unwrap() error {
	return e.Err
}

func main() {
	err := m()

	if err != nil {
		fmt.Println(err.Unwrap())
		return //если приложение полетит
	}

	fmt.Println("Success") //В случае nil error

}

func m() *customError {
	return &customError{
		Err:    fmt.Errorf("my error"),
		Custom: "baba",
	}
}
