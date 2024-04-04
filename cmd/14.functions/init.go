package functions

import "fmt"

func UseInit() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Saída: %d\n", i)
	}
}

func init() {
	fmt.Println("Função INIT")
}
