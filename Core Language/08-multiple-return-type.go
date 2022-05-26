package main

import "fmt"

func addSub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {

	add, _ := addSub(10, 20)

	fmt.Println("Sum: ", add)
}
