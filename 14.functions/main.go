package main

import "fmt"

func soma(l ...int) (total int) {
	for _, i := range l {
		total += i
	}
	return
}

func main() {
	fmt.Println(soma(1, 2, 3))
}
