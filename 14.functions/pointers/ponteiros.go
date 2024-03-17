package main

import "fmt"

func vezesTres(n *int) {
	*n = *n * 3
}

func main() {
	num := 10

	fmt.Println(num)
	vezesTres(&num)
	fmt.Println(num)
}
