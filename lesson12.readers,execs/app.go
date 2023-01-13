package main

import (
	"fmt"
	"os"
)

const namePath = "resources/test.txt"

func main() {
	AppendToFile("alex")
}

func ReadFile() {
	data, err := os.ReadFile(namePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func WriteFile(name string) {
	data := []byte(name)
	err := os.WriteFile(namePath, data, 0777)

	if err != nil {
		fmt.Println(err)
	}
}

func AppendToFile(name string) {
	f, err := os.OpenFile(namePath, os.O_APPEND|os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(name); err != nil {
		panic(err)
	}

}
