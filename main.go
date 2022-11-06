package main

import (
	"fmt"

	function "github.com/hadigun007/latihan-golang/04112022-function"
	structs "github.com/hadigun007/latihan-golang/04112022-struct"
	_ "github.com/hadigun007/latihan-golang/06112022-interface"
)

func main() {
	// functions()
	structs0()
}

func functions() {
	function.Function() // Latihan function

	var hasilAdd = function.Add(50, 20)
	fmt.Println(hasilAdd) // 70

	var hasilSub = function.Sub(10, 2)
	fmt.Println(hasilSub) // 8

	var a, b = function.MultVal(5, "value kedua")
	fmt.Println(a) // 5
	fmt.Println(b) // value kedua

	var ano = function.Ano("Hiii")
	fmt.Println(ano)
}

func structs0() {
	var ini = structs.StrctName{}
	ini.Name = "Hadi Gunawan"
	fmt.Println(ini.Name)

	var ini2 = structs.StrctName{}
	ini2.Name = "Tambunan"
	fmt.Println(ini2.Name)

	var ini3 = &structs.StrctName{Key: "1", Name: "Hadi", Age: 22, Gender: "lk"}
	fmt.Println(ini3.Age)

	var ini4 = &structs.StrctName{}
	ini4.Age = 29
	fmt.Println(ini4.Age)

}

func interface0() {

}
