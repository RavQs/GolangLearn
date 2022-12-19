package main

import "fmt"

func main() {
	fmt.Println("hello world")
	name := "babo"
	age := 20
	var c = fmt.Sprintf("My name is %s. And i'm %d years old", name, age)
	fmt.Println(c)
	//fmt.Println(test(20, 30))
	//deferLearn()
	pointersLearn()
}

func test(a, b int) int {
	return a + b
}

func calc(a, b int, math string) int {
	result := 0
	switch math {
	case "+":
		result = a + b
		break
	case "-":
		result = a - b
		break
	case "*":
		result = a * b
		break
	case "/":
		result = a / b
		break
	}
	return result
}

func cycleLearn() {
	i := 1
	for i < 1000 { //Cycle while
		if i == 100 {
			fmt.Println(fmt.Sprintf("cycle finally reached %d", i))
		}
		i++
	}

	fmt.Println(calc(5, 5, "*"))
}

func deferLearn() {
	defer fmt.Println("4")
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
	fmt.Println("start")
}

func pointersLearn() {
	a := "hello World"
	b := 42
	fmt.Println(a)
	fmt.Println(b)

	p := &a // где $ - адрес ячейки памяти ( амперсант)
	fmt.Println(p)
	fmt.Println(*p) // * - значение адреса ссылки
	*p = "oh my"
	fmt.Println(a) //значение a поменялось т.к. а ссылалось на ссылку а не на значение в памяти

}
