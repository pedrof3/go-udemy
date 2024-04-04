package main

import "fmt"

func vezesTres(n *int) {
	*n = *n * 3
}

func Usepointer() {
	num := 10

	fmt.Println(num)
	vezesTres(&num)
	fmt.Println(num)
}
