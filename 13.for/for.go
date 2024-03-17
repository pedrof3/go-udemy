package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("GO!!!")

	// Example 2
	slice := []string{"André", "Mila", "Paulo", "Pedro"}

	for _, nome := range slice {
		fmt.Printf("Hello, %s\n", nome)
	}
}
