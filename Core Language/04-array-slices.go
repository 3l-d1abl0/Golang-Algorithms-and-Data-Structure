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

	//SLICES - similar to arrays
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) //1 -2
	fmt.Println(fruitSlice[2])
	fmt.Println(fruitSlice[:4])

	//append is unique to slices.
	fruitSlice = append(fruitSlice, "Guava")

	fmt.Println(fruitSlice)

	fruitSlice = append(fruitSlice, "Berry", "Cherry")

	fmt.Println(fruitSlice)

	s := make([]int, 3)

	s[0] = 3
	s[1] = 2
	s[2] = 1

	fmt.Println(s)

	x := s //pointing to s
	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)

	//creat a new slice based on previous
	y := make([]int, len(s)+2)
	copy(y, s)

	y[0] = 1000
	fmt.Println(y)
	fmt.Println(s)

	//2D slices
	ss := make([][]int, 3)

	for i := 0; i < 3; i++ {

		ss[i] = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			ss[i][j] = i + j
		}
	}

	fmt.Println(ss)
}
