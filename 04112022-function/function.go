package function

import "fmt"

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
