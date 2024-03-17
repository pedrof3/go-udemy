package main

import "fmt"

func main() {
	names := []string{"jorge", "marcia", "carla", "augusto", "geovana"}
	names = append(names, "michael")

	fmt.Println(names)

	for _, name := range names {
		fmt.Println(name)
	}
}
