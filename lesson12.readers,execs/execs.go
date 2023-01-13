package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	c := exec.Command("systeminfo")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()

	time.Sleep(20 * time.Second)
}
