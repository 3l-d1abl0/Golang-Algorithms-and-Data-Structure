package main

import (
	"fmt"
	"time"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

	fmt.Println(1 + 1)

	fmt.Println(5 == 6)

	fmt.Println("1 + 1 is equal to ", 1+1)

	var price1, price2 float64 = 9.99, 9.95

	fmt.Println(price1, price2)

	fmt.Println("Time: ", time.Now())
}
