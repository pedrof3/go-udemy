package main

import "fmt"

func main() {
	num := [5]int{10, 20, 30, 40, 50}
	fmt.Println(num)

	num[0] = 15
	fmt.Println(num)
}
