package components

import (
	"fmt"
)

// Interfaces

type SomeStruct struct {
	N1, N2, N3 int
}

func (s *SomeStruct) Sum() int {
	fmt.Printf("N1: %d\nN2: %d\nN3: %d", s.N1, s.N2, s.N3)
	return s.N1 + s.N2 + s.N3
}

type SomeInterface interface {
	Sum() int
}
