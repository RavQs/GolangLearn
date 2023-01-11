package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go fmt.Println("World")
	fmt.Println("1")
	fmt.Println("2")
	go say(ch)

	for i := range ch {
		fmt.Println(i)
	}
	//Основной поток не завершится пока не прочтет данные из канала

	data := make(chan int)
	exit := make(chan int)

	go func() { //Анонимная функция
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()

	selectOne(data, exit)
}

func say(ch chan int) {
	time.Sleep(time.Second * 2)

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch) //close необходим т.к. возникнет deadlock (при итерации)
}

//selects

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x++
		case <-exit:
			fmt.Println("exit")
			return
		default:

		}

	}
}
