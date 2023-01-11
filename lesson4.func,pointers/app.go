package main

import "fmt"

//Callback functions,closures,pointers

type Point struct {
	X    int
	Y    int
	Name string
}

func (p *Point) testPointer(x, y int, name string) {
	p.X += x
	p.Y += y
	p.Name = name
}

func main() {
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	multiCallback := func(n1, n2 int) int {
		return n1 * n2
	}

	resultSum := doSomething(sumCallback, "something")
	resultMulti := doSomething(multiCallback, "something")

	fmt.Println(resultSum)
	fmt.Println(resultMulti)

	x := adder(50)
	fmt.Println(x(50))

	p := Point{X: 1, Y: 2}

	p.testPointer(5, 10, "bobo") //через указатель мы не создаем отдельный экземпляр структуры а ссылаемся на уже сущ-ую структуру

	fmt.Println(p)

}

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 2)
	fmt.Println(result)
	fmt.Println(s)

	return result
}

func adder(x int) func(int) int {
	sum := x
	return func(x int) int { //Замыкание (по сути static в java) они дешевле по производительности чем рекурсивные функции
		sum += x
		return sum
	}
}
