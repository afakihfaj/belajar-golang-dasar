package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	//pas by value
	// address1 := Address{"Situbondo", "Jawa Timur", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Bandung"
	// address2.Province = "Jawa Barat"
	// fmt.Println(address1) // tidak berubah apa2
	// fmt.Println(address2) // city dan province berubah menjadi bandung dan jawa barat
	// perubahan yang terjadi di address 2 ga ngaruh ke address 1

	//pass by reference
	var address1 Address = Address{"Situbondo", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Bandung"
	address2.Province = "Jawa Barat"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung,jawa barat

}
