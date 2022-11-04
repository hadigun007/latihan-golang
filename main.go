package main

import (
	"fmt"

	function "github.com/hadigun007/latihan-golang/04112022-function"
)

func main() {
	functions()
}

func functions() {
	function.Function() // Latihan function
	var hasilAdd = function.Add(50, 20)
	fmt.Println(hasilAdd) // 70
	var hasilSub = function.Sub(10, 2)
	fmt.Println(hasilSub) // 8
	var a, b = function.MultVal(5, "value kedua")
	fmt.Println(a) // 5
	fmt.Println(b) // value
	var age = 20
	var text = "John"
	fmt.Println("Before : ", age, text) // Before : 5 John
	function.Update(&age, &text)
	fmt.Println("after : ", age, text) // After : 20 John John
}
