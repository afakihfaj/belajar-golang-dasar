package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var afakih Customer
	afakih.Name = "Afakih Fajduwani"
	afakih.Address = "Indonesia"
	afakih.Age = 26

	fmt.Println(afakih)
	fmt.Println(afakih.Name)
	fmt.Println(afakih.Address)
	fmt.Println(afakih.Age)

	laras := Customer{
		Name:    "Laras",
		Address: "Indonesia",
		Age:     26,
	}
	fmt.Println(laras)

	ayudia := Customer{"Ayudia", "Indonesia", 26}
	fmt.Println(ayudia)

	ayudia.sayHello("Baek Chanbi")
	laras.sayHello("Baek Chanbi")
	afakih.sayHello("Baek Chanbi")
}
