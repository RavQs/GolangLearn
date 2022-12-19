package main

import "fmt"

//arrays,slices

func main() {
	//arrLearn()
	slicesLearn()
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
	//если len > cap то cap увеличивается вдвое

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
