package main

import (
	"fmt"
)

func problemsolvingtest3() string {
	var arr, leftsum, rightsum int

	fmt.Println("Array Size:")
	fmt.Scanln(&arr)

	numbers := make([]int, arr)
	for i := 0; i < arr; i++ {
		fmt.Scanln(&numbers[i])
	}

	for i := 1; i < arr; i++ {
		for j := i - 1; j >= 0; j-- {
			leftsum += numbers[j]
		}
		for k := i + 1; k < arr; k++ {
			rightsum += numbers[k]
		}
		if leftsum == rightsum {
			return "yes"
		}
	}
	return "no"
}
