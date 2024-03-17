package main

import (
	"fmt"
)

func fullName() (string, string) {
	firstName := "Pedro"
	lastName := "Ferreira"
	return firstName, lastName
}

func main() {
	calculando := func(a, b int) (int, int) {
		soma := a + b
		multiplicacao := a * b
		return soma, multiplicacao
	}

	fmt.Println(calculando(21, 22))
	nome, _ := fullName()
	_, sobrenome := fullName()

	fmt.Println(nome)
	fmt.Println(sobrenome)
}
