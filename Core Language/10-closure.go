package main

import "fmt"

func regularFunction(msg string) {
	fmt.Println(msg)
}

//Closure - returns a function with argument type string
func return_anon_func() func(string) {

	//returns an anonymous function
	return func(msg string) {
		fmt.Println(msg)
	}
}

//Closure example
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//1. Regular function
	regularFunction("Hello")

	//2. Anonymous Function - A function has no Name
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymous")

	//3. Closures
	print_func := return_anon_func()
	print_func("Hello returned from anonymous")

	//4. Example
	next_int := int_seq()
	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())
}
