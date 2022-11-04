package function

import "fmt"

// Function name must be in Caapitalize tobe visible
func Hello() {
	fmt.Println("Hello World")
}

func Add(x, y int) int {
	total := 0
	total = x + y
	return total
}
