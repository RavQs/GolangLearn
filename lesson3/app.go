package main

import "fmt"

//Iterators,map structure

type Point struct {
	X    int
	Y    int
	Name string
}

func (p Point) method() { //метод структуры
	fmt.Println(p)
}

func main() {
	//forRange()
	mapLearn()
}

func forRange() {
	arr := []string{"a", "b", "c"}
	for i, l := range arr {
		fmt.Println(fmt.Sprintf("%d : %s", i, l))
	}

	for _, l := range arr { // для игнора индекса
		fmt.Println(l)
	}
}

func mapLearn() {
	pointMap := map[string]Point{}
	pointMap["a"] = Point{X: 1, Y: 2, Name: "babo"}
	pointMap["b"] = Point{X: 2, Y: 3, Name: "kimSSSs"}

	otherMap := map[string]int{}
	otherMap["a"] = 5
	otherMap["b"] = 10
	otherMap["c"] = 15

	var oneMorePointMap = map[string]Point{
		"a": {X: 15, Y: 20},
		"b": {X: 50, Y: 100},
	}

	printPoint(oneMorePointMap)

	value, ok := oneMorePointMap["a"] //Проверка ключа в мапе

	if ok {
		fmt.Println(value)
	}
}

func printPoint(point map[string]Point) {
	fmt.Println(point)
	point["a"].method()
	point["b"].method()

}
