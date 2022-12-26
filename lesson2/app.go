package main

import "fmt"

//arrays,slices

func main() {
	//arrLearn()
	//slicesLearn()
	slicesLearn2()
}

func arrLearn() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[0])

	nums := [...]int{1, 2, 3, 4} //Подсчет массива
	// ... = компилятор сам подсчитывает размерность
	nums[3] = 4

}

func slicesLearn() {
	//Нет фиксированной длины
	//Слайсы это массивы с методами
	//1 вызов append создает копию слайса с capacity в 2 раза больше(но лишь 1 раз)
	//и копирует туда старый с добавленным значением(ями). С каждым новым присваиванием функции, если выйти за пределы вместимости, то тенденция мултиплаинга повторится

	//Можно создать слайс с фиксированной длиной
	createSlice := make([]string, 3)
	createSlice[0] = "I"
	createSlice[1] = "II"
	createSlice[2] = "III"
	//createSlice[3] = "IV" //Компилятор ничего не говорит но ошибка будет
	//необходимо использовать append
	createSlice = append(createSlice, "IV")
	fmt.Println(createSlice)
}

func slicesLearn2() {
	animalArr := [4]string{
		"cat",
		"dog",
		"giraffe",
		"bird",
	}
	fmt.Println(animalArr)
	a := animalArr[0:2] //слайс с 2 length
	fmt.Println(a)
	b := animalArr[2:4]
	fmt.Println(b)
	c := animalArr[1:3]
	fmt.Println(c)
	d := animalArr[:3] //c нуля до 3
	e := animalArr[2:] //с 2 до конца
	f := animalArr[:]  //с нуля и до конца
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
