package main

import (
	"fmt"
	"modulos/adress"
)

func main() {
	endereco := "rua das palemiras"
	fmt.Println(adress.TipoDeEndereco(endereco))
}
