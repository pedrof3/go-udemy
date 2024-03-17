package main

import "fmt"

type person struct {
	name string
	age  int
}

type player struct {
	person
	club     string
	position string
}

func main() {
	// Método 1: Utilizando um tipo de dados personalizado STRUCT
	var p1 person
	p1.name = "Pedro"
	p1.age = 20

	// Método 2: Utilizando um tipo de dados personalizado STRUCT
	p2 := person{"Maria", 21}

	// Método 1: Utilizando um tipo de dados personalizado com "HERANÇA"
	var j1 player
	j1.name = "Messi"
	j1.age = 28
	j1.club = "Barcelona"
	j1.position = "Striker"

	// Método 2: Utilizando um tipo de dados personalizado com "HERANÇA"
	j2 := player{person{"Pelé", 17}, "Santos", "Striker"}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(j1)
	fmt.Println(j2.name)
}
