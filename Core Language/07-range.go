package main

import "fmt"

func main() {

	//1.Over Arrays
	str_arr := []string{"a", "b", "c", "d"}

	for idx, value := range str_arr {

		fmt.Printf("idx: %d val: %s \n", idx, value)
	}

	//2. Over Maps
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println("Key: ", k, "val: ", v)
	}

	//3. Just Keys
	for k := range m {
		fmt.Println("Key: ", k)
	}

}
