package main

import "fmt"

func main() {
	name := "Afakih"

	switch name {
	case "Afakih":
		fmt.Println("Halo Afakih")
	case "Eko":
		fmt.Println("Halo Eko")
	case "Joko":
		fmt.Println("Halo Joko")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	} //switch tanpa stetement

	switch length := len(name); length > 6 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	} //switch dengan stetement
	name = "JamalismJamalism"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 6:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")

	}
}
