package main

import (
	"fmt"
	"reflect"
)

// Interfaces

type SomeStruct struct {
	N1, N2, N3 int
}

func (s *SomeStruct) Sum() int {
	fmt.Println(fmt.Sprintf("N1: %d\nN2: %d\nN3: %d", s.N1, s.N2, s.N3))
	return s.N1 + s.N2 + s.N3
}

type SomeInterface interface {
	Sum() int
}

func main() {
	s := SomeStruct{N1: 1, N2: 5, N3: 15}

	if reflect.TypeOf(s).Name() == "SomeStruct" {
		fmt.Println("typeof in Golang is working")
	}

	fmt.Println(s.Sum())
}
