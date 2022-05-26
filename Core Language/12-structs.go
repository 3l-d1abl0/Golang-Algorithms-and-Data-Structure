package main

import "fmt"

type employee struct {
	first_name string
	last_name  string
	id         int
}

func main() {

	fmt.Println(employee{
		first_name: "Frank",
		last_name:  "Castle",
		id:         1})

	fmt.Println(employee{
		"Matthew",
		"Murdock",
		2})

	fmt.Println(employee{
		first_name: "Foggy",
		last_name:  "Nelson"})

	emp := employee{
		first_name: "Frank",
		last_name:  "Castle",
		id:         1}

	fmt.Println(emp.first_name)

	//Pointers to employee
	emp_ptr := &emp

	emp_ptr.first_name = "Clark"
	fmt.Println(emp)
	fmt.Println(emp_ptr)
}
