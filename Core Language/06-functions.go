package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

//both num1, num2 is integer
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(getSum(3, 4))

	fmt.Println(greeting("YeeeHaaaw"))
}
