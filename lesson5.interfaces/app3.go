package main

import "fmt"

//interface, casting, type assertions, type switch

func main() {
	var a interface{} = "hello"
	//interface может быть типом данных
	//интерфейс используются как элемент слабой типизации
	//пустой интерфейс может хранить в себе элемент любого типа
	//println в аргументах содержит интерфейс(т.е. любой тип данных)

	//fmt.Println(a)
	//fmt.Println(fmt.Sprintf("%v : %T",a,a))
	//
	//a = 42
	//fmt.Println(a)
	//fmt.Println(fmt.Sprintf("%v : %T",a,a))

	s := a.(string)
	fmt.Println(s)

	f, ok := a.(float32) //где ok это успешность кастинга
	fmt.Println(f, ok)

	//type switch для проверки нескольких типов

	switch a.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case float32:
		fmt.Println("type is float32")
	case bool:
		fmt.Println("type is bool")
	default:
		fmt.Println("type is unknown")

	}

}
