package functions

import "fmt"

func RecursiveFibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return RecursiveFibonacci(n-2) + RecursiveFibonacci(n-1)
}

func UseFibonacci() {
	num := uint(45)

	for i := uint(1); i < num; i++ {
		fmt.Println(RecursiveFibonacci(i))
	}
}
