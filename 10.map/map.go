package main

import "fmt"

func main() {
	meuMap := map[string]int{
		"idade":   21,
		"salario": 2500,
	}
	fmt.Println(meuMap["salario"])

	mapAninhado := map[string]map[string]string{
		"carro": {
			"marca":  "chevrolet",
			"modelo": "celta",
			"cor":    "azul",
			"placa":  "a2b1c3",
		},
	}
	fmt.Println(mapAninhado["carro"]["placa"])

	delete(mapAninhado["carro"], "cor")

	fmt.Println(mapAninhado)
}
