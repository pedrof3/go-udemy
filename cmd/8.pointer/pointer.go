package pointers

import "fmt"

func sumTen(num *int) {
	*num = *num + 10
}

func main() {
	// Um ponteiro é uma variável que aponta o local da memória em que outra variável está armazenada
	nome := "Bruno"
	p := &nome

	fmt.Println(p)
}
