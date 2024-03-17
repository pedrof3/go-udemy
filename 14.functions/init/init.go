package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Saída: %d\n", i)
	}
}

func init() {
	fmt.Println("Função INIT")
}
