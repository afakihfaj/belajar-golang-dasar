package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Afakih"
	middleName = "Fajduwani"
	lastName = "Goat"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
