package main

import "fmt"

func main() {
	name := "Afakih"

	if name == "Afakih" {
		fmt.Println("Halo Afakih")
	} else if name == "Laras" {
		fmt.Println("Halo Larassssssss")
	} else if name == "Diana" {
		fmt.Println("Halo Dianaa")
	} else if name == "Anna" {
		fmt.Println("Halo Annaa")
	} else {
		fmt.Println("Halo, Boleh Kenalan Kah Manies?")
	}

	if length := len(name); length > 6 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	} //if short statement ya
}
