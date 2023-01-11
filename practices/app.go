package main

import (
	"fmt"
	"github.com/RavQs/GolangLearn/practices/myPractice"
)

func main() {
	arr := []int{0, 1, 0, 1, 0}
	fmt.Println(myPractice.FindOdd(arr))
	fmt.Println(myPractice.Convert("aabbccc"))

	u := myPractice.User{
		Id:   1,
		Name: "Babo",
		Pw:   "123123",
	}

	fmt.Println(u)

	arr2 := []int{5, 7, 234, 5, 33, -15}
	min, max := myPractice.FindMaxMin(arr2)
	fmt.Println(fmt.Sprintf("min value is %d\nmax value is %d", min, max))

}
