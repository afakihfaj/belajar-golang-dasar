package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	afakih := Man{"Afakih"}
	afakih.Married()

	fmt.Println(afakih.Name)
}
