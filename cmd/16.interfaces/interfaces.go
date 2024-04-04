package interfaces

import "fmt"

type acao interface {
	som() string
}

type cachorro struct {
	nome  string
	raca  string
	idade uint8
}

func (c cachorro) som() string {
	return "auauau"
}

type gato struct {
	nome         string
	fiosNoBigode int
}

func (g gato) som() string {
	return "miau"
}

func main() {
	c1 := cachorro{
		"bidu",
		"vira-lata",
		5,
	}
	g1 := gato{
		"mel",
		3500,
	}

	fmt.Println(c1.som())
	fmt.Println(g1.som())
}
