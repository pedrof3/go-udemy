package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type arena struct {
	Nome        string `json: nome`
	Time        string `json: string`
	Localidade  string `json: localidade`
	Capacidade  uint   `json: capacidade`
	Inauguracao uint   `json: inauguracao`
}

func main() {
	arenaCorinthians := arena{"Arena Corinthians", "Corinthians", "São Paulo", 43500, 2014}
	fmt.Println(arenaCorinthians)

	arenaCorinthiansJSON, err := json.Marshal(arenaCorinthians)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arenaCorinthiansJSON)
	fmt.Println(bytes.NewBuffer(arenaCorinthiansJSON))
}
