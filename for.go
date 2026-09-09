package main

import (
	"fmt"
)

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke", counter)
	// 	counter++
	// }
	// fmt.Println("Selesai")
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke", counter)
	}
	fmt.Println("Selesai") //for biasa

	names := []string{"Afakih", "Fajduwani", "Goat"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	} //for statement

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	} //range for tipe 1
	for _, name := range names {
		fmt.Println(name)
	} //range for tipe 2
}
