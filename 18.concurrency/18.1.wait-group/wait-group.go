package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("Programa terminado")
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		escrever("2012 CORINTHIANS")
		wg.Done()
	}()

	go func() {
		escrever("2003 PEDRO")
		wg.Done()
	}()
	wg.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
