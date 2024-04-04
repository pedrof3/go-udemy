package functions

import "fmt"

func VariadicSum(l ...int) (total int) {
	for _, i := range l {
		total += i
	}
	return
}

func UseVariadicSum() {
	fmt.Println(VariadicSum(1, 2, 3))
}
