package functions

import "fmt"

func sum(num ...int) (total int) {
	total = 0
	for _, valor := range num {
		total += valor
	}
	return
}

func AnonymousFunction() {
	calculando := sum(1, 43, 43645, 63463, 366)
	fmt.Println(calculando)

	// Função anônima
	var idade int = 12
	func(i int) {
		switch {
		case i < 18:
			fmt.Println("Menor de idade")
		default:
			fmt.Println("Maior de idade")
		}
	}(idade)
}
