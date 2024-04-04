package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "GoLang"
	canal <- "I'm learning GO!!!"

	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
