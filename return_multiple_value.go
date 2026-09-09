package main

import "fmt"

func getFullName() (string, string, string) {
	return "Afakih", "Fajduwani", "Goat"
}

func main() {
	// firstName, lastName, goatName := getFullName()
	// fmt.Println(firstName, lastName, goatName)
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
}
