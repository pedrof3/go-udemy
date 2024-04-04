package main

import (
	"fmt"
	"time"
)

/*
Tendo dois ou mais canais paralelos em Go, usamos o select para garantir que funcionem de maneira independente,nesse exemplo,
sem o uso do select, a função que mais rápida precisaria esperar a execução da função mais longa para poder ser executada novamente,
com isso as nossas goroutines não teriam o efeito desejado de paralelismo
*/
func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			canal1 <- "11 ----- 11"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			canal2 <- "22 ***** 22"
		}
	}()

	for {
		select {
		case msgCanal1 := <-canal1:
			fmt.Println(msgCanal1)

		case msgCanal2 := <-canal2:
			fmt.Println(msgCanal2)
		}
	}
}
