package main

import (
	"errors"
	"fmt"
	"log"
)

type AppError struct {
	Message string
	Err     error
}

func (ae AppError) Error() string {
	return ae.Message
}

func main() {
	divideByZero(5, 0)
	fmt.Println("hello")

}

func divideByZero(a, b int) int {
	var appError *AppError

	defer func() {
		if err := recover(); err != nil { //recover не роняет приложение в случае паники
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appError) {
					fmt.Println("panic!", err)
				} else {
					panic("custom panic")
				}
			default:
				panic("some default panic")
			}
			log.Println("panic happened", err)
		}
	}()
	return a / b
}

func div(a, b int) int {
	if b == 0 {
		panic(&AppError{
			Message: "Custom Error",
			Err:     nil})
	}

	return a / b
}

func baboTest() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("babo detected but its ok")
		}
	}()
	panic("babo is detected") //Можно указать панику вручную
}
