package main

import "fmt"

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	num := uint(45)

	for i := uint(1); i < num; i++ {
		fmt.Println(fibonacci(i))
	}
}
