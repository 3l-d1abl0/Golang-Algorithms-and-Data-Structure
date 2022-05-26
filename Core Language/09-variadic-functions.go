package main

import "fmt"

func mult(nums ...int) int {
	total := 1

	for _, num := range nums {
		total *= num
	}

	return total
}

func main() {

	//eg. Println
	fmt.Println("this", "is", "a", "variadic", "function")

	//call with multiple variables
	fmt.Println(mult(1, 2, 3, 4))

	//can be called with Slices
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(mult(nums...))
}
