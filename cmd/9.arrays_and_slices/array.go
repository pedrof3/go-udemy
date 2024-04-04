package arrays_and_slices

import "fmt"

func BasicArrays() {
	nums := [5]int{10, 20, 30, 40, 50}
	fmt.Println(nums)

	var names [6]string

	for i := 0; i < len(names); i++ {
		var newName string
		fmt.Print("Digite o 1° nome: ")
	}

	nums[0] = 15
	fmt.Println(nums)
}
