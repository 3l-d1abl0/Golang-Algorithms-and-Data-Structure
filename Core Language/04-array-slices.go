package main

import "fmt"

func main() {
	var intArr [10]int
	fmt.Println(intArr) //default value 0

	var boolArr [10]bool
	fmt.Println(boolArr) //default value false

	var fruitArr [5]string
	fmt.Println(fruitArr)

	// // Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fruitArr[4] = "Grapes"
	fmt.Println(fruitArr)

	// Decalre and assign
	fruits := [2]string{"Apple", "Orange"}

	fmt.Println(fruits)
	fmt.Println(fruits[1])

	var aa [5][5]int
	fmt.Println(aa)

	count := 1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			aa[i][j] = count
			count = count + 1
		}
	}

	fmt.Println(aa)

	//SLICES
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
