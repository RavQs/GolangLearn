package main

import "fmt"

//vars,defers,cycles,pointers,structs

func main() {
	//fmt.Println(test(20, 30))
	//varLearn()
	//deferLearn()
	//pointersLearn()
	//structLearn()

	p1 := Point{X: 1,
		Y:    2,
		Name: "babo",
	}
	p1.PointPrintMethod()
}

func test(a, b int) int {
	return a + b
}

func varLearn() {
	fmt.Println("hello world")
	name := "babo"
	age := 20
	var c = fmt.Sprintf("My name is %s. And i'm %d years old", name, age)
	fmt.Println(c)
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

type Point struct {
	X    int
	Y    int
	Name string
}

func (p *Point) PointPrintMethod() { //Добавление метода к структуре (метод у класса)
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.Name)
}

func structLearn() {
	p1 := Point{X: 1, Y: 2}
	fmt.Println(p1)
	fmt.Println(p1.Y)
	p1.Y = 15
	fmt.Println(p1.Y)
}
