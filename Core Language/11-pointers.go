package main

import "fmt"

func main() {

	val := 20
	fmt.Println(val)

	//1. address of val
	ptr := &val

	fmt.Println(ptr, val)

	fmt.Printf("%T %T\n", ptr, val)

	//2. Dereferencing
	fmt.Println(*ptr)

	*ptr = 10

	fmt.Println(*ptr, val)
	fmt.Printf("%T %T\n", *ptr, val)

	val = 50
	fmt.Println(*ptr)

}
