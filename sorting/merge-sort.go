package main

import "fmt"

func _merge(arr []int, left int, mid int, right int) {

	//sdsdsd
	//sdjkbsdvk
	//skjdbdsjbv
	leftArray := arr[left:mid]
	rightArray := arr[mid:right]

	i, j, k := 0, 0, left

	fmt.Println(i, j, k)
	fmt.Println(leftArray)
	fmt.Println(rightArray)

	fmt.Println(arr)

	for i < len(leftArray) && j < len(rightArray) {

		fmt.Println(i, j, k)

		if leftArray[i] <= rightArray[j] {
			arr[k] = leftArray[i]
			i++
		} else {
			arr[k] = rightArray[j]
			j++
		}

		k++
	}

	for i < len(leftArray) {
		arr[k] = leftArray[i]
		i++
		k++
	}

	for j < len(rightArray) {
		arr[k] = rightArray[j]
		j++
		k++
	}

	fmt.Println(arr)
	fmt.Println("========")

}

func mergeSort(arr []int, left int, right int) {

	if left < right { //len==1
		var mid int = left + (right-left)/2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		_merge(arr, left, mid, right)
	}
}

func main() {

	arr := []int{2, 4, 3, 1, 6, 8, 4}

	mergeSort(arr, 0, len(arr))

	fmt.Println(arr)
}
