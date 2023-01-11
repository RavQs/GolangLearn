package myPractice

import "fmt"

func FindOutlier(integers []int) int {
	fakeOddNum := 0
	fakeEvenNum := 0
	counter := 0

	for _, l := range integers {
		if l%2 == 0 || l == 2 {
			fakeOddNum = l
			counter++
		} else {
			fakeEvenNum = l
			counter--
		}
	}

	if counter > 0 {
		return fakeEvenNum
	}
	return fakeOddNum
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

func DigitalRoot(n int) int {
	// Если число больше 100 число отсылается в метод через цикл
	// 16
	// Если число < 10 return
	tempNum := 0
	result := 0
	for n != 0 {
		tempNum = n % 10
		result += tempNum
		n = n / 10
		fmt.Println(n)
	}

	if result >= 10 {
		result = DigitalRoot(result)
	}

	return result
}
