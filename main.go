package main

import (
	"fmt"
	"time"
)

func do(i int) {
	fmt.Println(i)
}

func main() {
	msg := "Let's Go"
	go func(input string) {
		fmt.Println(input)
	}(msg)
	msg = "Let's Go Go Go"
	time.Sleep(1 * time.Second)
}
