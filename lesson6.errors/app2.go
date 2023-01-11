package main

import (
	"fmt"
	"io"
	"strings"
)

//Reader

func main() {
	//io.Reader() // интерфейс чтения файла
	//io.EOF //конец файла( в голанге работает как ошибка)
	r := strings.NewReader("hello world")

	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n: %d error: %v byte: %v\n", n, err, arr)
		fmt.Printf("arr in bytes: %q", arr[:n])
		if err == io.EOF {
			break
		}
	}
}
