package main

import (
	"fmt"
	"time"
)

func mult(n int, canal chan int) {
	for i := 0; i < 10; i++ {
		canal <- i * n
		time.Sleep(time.Millisecond * 200)
	}
	close(canal)
}

func main() {
	canal := make(chan int)

	go mult(3, canal)
	go mult(21, canal)

	for message := range canal {
		fmt.Println(message)
	}
}
