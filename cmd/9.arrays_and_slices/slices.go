package arrays_and_slices

import "fmt"

func BasicSlices() {
	names := []string{"jorge", "marcia", "carla", "augusto", "geovana"}
	names = append(names, "michael")

	fmt.Println(names)

	for _, name := range names {
		fmt.Println(name)
	}
}
