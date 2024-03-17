package main

import "fmt"

// defer executa um bloco de código somente no final de função que está inserida
func main() {
	defer fmt.Println("Teste 1")
	fmt.Println("Teste 2")
	fmt.Println("Teste 3")
}
