package main

import (
	"fmt"
)

func f(a []int) {
	//fmt.Println(a[1])
	a[5] = 67
	//fmt.Println(a)
}

func main() {
	a := make([]int, 10, 100)
	a[1] = 100
	fmt.Println(a)
	f(a)

	fmt.Println(a)
}
