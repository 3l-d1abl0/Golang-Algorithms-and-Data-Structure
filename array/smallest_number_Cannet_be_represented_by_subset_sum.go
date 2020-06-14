package main

import "fmt"

//https://www.geeksforgeeks.org/find-smallest-value-represented-sum-subset-given-array/

func findSmallest(arr []int) int {

	maxPoss := 0

	if len(arr) == 0 || arr[0] != 1 {
		return 1
	}

	maxPoss = 1

	for idx := 1; idx < len(arr); idx++ {

		if arr[idx] > maxPoss+1 {
			break
		}

		maxPoss += arr[idx]
	}

	return maxPoss + 1

}

func main() {

	arr := []int{1, 1, 3, 5, 8, 21}
	fmt.Println(findSmallest(arr))
	fmt.Println()

	arr = []int{1, 3, 4, 5}
	fmt.Println(findSmallest(arr))
	fmt.Println()

	arr = []int{1, 2, 6, 10, 11, 15}
	fmt.Println(findSmallest(arr))
	fmt.Println()

	arr = []int{1, 1, 1, 1}
	fmt.Println(findSmallest(arr))
	fmt.Println()

	arr = []int{1, 1, 3, 4}
	fmt.Println(findSmallest(arr))
	fmt.Println()

}
