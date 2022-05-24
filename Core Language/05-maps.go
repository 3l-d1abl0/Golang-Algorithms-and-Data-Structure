package main

import "fmt"

func main() {

	//Maps Method 1
	m := make(map[string]int)

	m["a"] = 0
	m["b"] = 1
	m["c"] = 2

	fmt.Println(m)
	fmt.Println(m["b"])
	fmt.Println(len(m))

	delete(m, "b")
	fmt.Println(m)

	//Maps Method 2
	// Decalre map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	val, is_present := m["Z"]
	fmt.Println(is_present, val)

	_, present := m["b"]
	fmt.Println(present)
}
