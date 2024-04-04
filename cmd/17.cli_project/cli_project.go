package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pedrof3/go-udemy/cmd/17.cli_project/app"
)

func main() {
	fmt.Println("Programa CLI")
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
