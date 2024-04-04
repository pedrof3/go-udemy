package main

import (
	"fmt"
	"log"
	"os"
	"program/app"
)

func main() {
	fmt.Println("Programa CLI")
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
