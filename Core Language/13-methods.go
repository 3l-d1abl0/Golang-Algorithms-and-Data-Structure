package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {

	r := rect{width: 10, height: 5}
	fmt.Println("Area : ", r.area())

	r_ptr := &r
	fmt.Println("Perimeter : ", r_ptr.perim())
}
