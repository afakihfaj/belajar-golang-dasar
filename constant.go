package main

import "fmt"

func main() {
	const (
		firstName string = "Afakih"
		lastName         = "Fajduwani"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "koko"
}
