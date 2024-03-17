// Worker pools pattern
package main

import "fmt"

func main() {
	fmt.Println("Concurrency patterns: Worker pools")
	tarefas := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tarefas, results)
	go worker(tarefas, results)
	go worker(tarefas, results)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(tarefas <-chan int, results chan<- int) {
	for num := range tarefas {
		results <- fibonacci(num)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
