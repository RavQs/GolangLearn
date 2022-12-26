package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 1, 0}

	fmt.Println(FindOdd(arr))
}

func FindOdd(seq []int) int {
	counter := 0

	for i := 0; i < len(seq); i++ {
		for j := 0; j < len(seq); j++ {
			if seq[j] == seq[i] {
				counter++
			}
		}
		if counter%2 != 0 {
			return seq[i]
		}
	}
	return counter
}
