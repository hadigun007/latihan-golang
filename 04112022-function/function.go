package function

import "fmt"

// learning resource
// https://www.golangprograms.com/go-language/functions.html

// Function name must be in Caapitalize tobe visible
func Function() {
	fmt.Println("Latihan function")
}

func Sub(x, y int) int {
	return x - y

}
func Add(x, y int) (hasil int) {
	hasil = x + y
	return
}

// return multiple value
func MultVal(pertama int, kedua string) (prt int, kda string) {
	prt = pertama
	kda = kedua
	return
}

// passing address to a function
// override given data
func Update(a *int, t *string) {
	*a = *a + 5
	*t = *t + " John"
	return
}
