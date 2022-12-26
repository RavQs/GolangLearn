package main

import "fmt"

func main() {
	//fake := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	//fmt.Println(FindOutlier(fake))

	//fmt.Println(InAscOrder([]int{21, 12, 4, 7, 19}))

	fmt.Println(Gimme([3]int{1, 2, 3}))
}

func FindOutlier(integers []int) int {
	arr2 := make([]string, 3)
	fakeOddNum := 0
	fakeEvenNum := 0
	counter := 0

	for _, l := range integers {
		if l%2 == 0 || l == 2 {
			arr2 = append(arr2, "even")
			fakeOddNum = l
			counter++
		} else {
			arr2 = append(arr2, "odd")
			fakeEvenNum = l
			counter--
		}
	}

	if counter > 0 {
		return fakeEvenNum
	} else {
		return fakeOddNum
	}
}

func InAscOrder(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return false
		}
	}
	return true
}

func Gimme(array [3]int) int {
	min := array[0]
	max := array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	for item, value := range array {
		if value != min && value != max {
			return item
		}
	}

	return 0

}
