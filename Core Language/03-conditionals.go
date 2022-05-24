package main

import "fmt"

func main() {
	x := 10
	y := 10

	//  If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT blue or red")
	}

	q := 10

	switch q {
	case 1, 2:
		fmt.Println("One or Two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Default")
	}

	k := 100
	switch {
	case q == 10:
		fmt.Println(q, " is equal to 10")

	case k > 90:
		fmt.Println(k, "greter than 90")
	default:
		fmt.Println("Executing Default")
	}
}
