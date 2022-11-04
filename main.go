package main

import (
	"fmt"

	function "github.com/hadigun007/latihan-golang/04112022-function"
	structs "github.com/hadigun007/latihan-golang/04112022-struct"
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
}
