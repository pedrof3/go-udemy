package main

import "fmt"

func main() {
	heights := make([]int, 5)

	fmt.Println(heights)

	for i := 0; i < len(heights); i++ {
		heights[i] = i * 2
	}

	fmt.Println(heights)
}
