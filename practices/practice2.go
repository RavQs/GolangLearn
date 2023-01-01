package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//arr := []int{0, 1, 0, 1, 0}
	//fmt.Println(FindOdd(arr))

	fmt.Println(Convert("abbbbbbbbbbbbcccccccccccccddddddrrrrrrrrrrrrrrrrrrrffffffffeRRRRRRRRR"))
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

func Grow(arr []int) int {
	result := 1

	for i := 0; i < len(arr); i++ {
		result *= arr[i]
	}
	return result
}

func SubtractSum(n int) string {
	var fruits = map[int]string{
		1:   "kiwi",
		2:   "pear",
		3:   "kiwi",
		4:   "banana",
		5:   "melon",
		6:   "banana",
		7:   "melon",
		8:   "pineapple",
		9:   "apple",
		10:  "apple",
		11:  "cucumber",
		12:  "pineapple",
		13:  "cucumber",
		14:  "orange",
		15:  "grape",
		16:  "orange",
		17:  "grape",
		18:  "apple",
		19:  "grape",
		20:  "cherry",
		21:  "pear",
		22:  "cherry",
		23:  "pear",
		24:  "kiwi",
		25:  "banana",
		26:  "kiwi",
		27:  "apple",
		28:  "melon",
		29:  "banana",
		30:  "melon",
		31:  "apple",
		32:  "melon",
		33:  "pineapple",
		34:  "cucumber",
		35:  "orange",
		36:  "apple",
		37:  "orange",
		38:  "grape",
		39:  "orange",
		40:  "grape",
		41:  "cherry",
		42:  "pear",
		43:  "cherry",
		44:  "pear",
		45:  "apple",
		46:  "pear",
		47:  "kiwi",
		48:  "banana",
		49:  "kiwi",
		50:  "banana",
		51:  "melon",
		52:  "pineapple",
		53:  "melon",
		54:  "apple",
		55:  "cucumber",
		56:  "pineapple",
		57:  "cucumber",
		58:  "orange",
		59:  "cucumber",
		60:  "orange",
		61:  "grape",
		62:  "cherry",
		63:  "apple",
		64:  "cherry",
		65:  "pear",
		66:  "cherry",
		67:  "pear",
		68:  "kiwi",
		69:  "pear",
		70:  "kiwi",
		71:  "banana",
		72:  "apple",
		73:  "banana",
		74:  "melon",
		75:  "pineapple",
		76:  "melon",
		77:  "pineapple",
		78:  "cucumber",
		79:  "pineapple",
		80:  "cucumber",
		81:  "apple",
		82:  "grape",
		83:  "orange",
		84:  "grape",
		85:  "cherry",
		86:  "grape",
		87:  "cherry",
		88:  "pear",
		89:  "cherry",
		90:  "apple",
		91:  "kiwi",
		92:  "banana",
		93:  "kiwi",
		94:  "banana",
		95:  "melon",
		96:  "banana",
		97:  "melon",
		98:  "pineapple",
		99:  "apple",
		100: "pineapple",
	}

	for n > 100 {
		n -= getSum(n)
	}
	fmt.Println(n)
	return fruits[n]
}

func getSum(n int) int {
	if n < 10 {
		return n
	}
	return getSum(n/10) + n%10
}

func QuarterOf(month int) int {
	if month >= 1 && month <= 3 {
		return 1
	} else if month >= 4 && month <= 6 {
		return 2
	} else if month >= 7 && month <= 9 {
		return 3
	}
	return 4
}

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func NoSpace(word string) string {
	word = strings.ReplaceAll(word, " ", "")
	return strings.TrimSpace(word)
}

//"abbbcc" -> "a1bbb3cc2"

func Convert(name string) string {
	counter := 0
	result := ""
	for i := 0; i < len(name); i++ {
		counter = strings.Count(name, string(name[i]))

		result += string(name[i])
		if i == len(name)-1 || name[i] != name[i+1] {
			result += strconv.Itoa(counter)
		}
	}

	return result
}
