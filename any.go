package main

import "fmt"

func Ups() any {
	//return 1
	//return true
	return "Ups"

	//tipe data yang paling tinggi posisinya,semuanya mengikuti interface kosong/any
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)

}
