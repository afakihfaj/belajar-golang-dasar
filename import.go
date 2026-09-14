package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Afakih")
	fmt.Println(result)
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Afakih")) // tidak bisa diakses karena dari package yang berbeda dan namanya pake huruf kecil
}
