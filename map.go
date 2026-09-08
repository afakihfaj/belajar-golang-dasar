package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Afakih"
	// person["address"] = "Situbondo"
	person := map[string]string{
		"name":    "Afakih",
		"address": "Situbondo",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Afakih"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
