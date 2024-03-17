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

func multiplexing(entrada1, entrada2 <-chan string) <-chan string {
	saida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-entrada1:
				saida <- mensagem
			case mensagem := <-entrada2:
				saida <- mensagem
			}
		}
	}()

	return saida
}

func main() {
	canal := multiplexing(escrever("Olá mundo"), escrever("Olá golang"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
