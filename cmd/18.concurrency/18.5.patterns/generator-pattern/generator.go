package main

import (
	"fmt"
	"time"
)

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("output: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}

func main() {
	canal := escrever("Minha vez...")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
