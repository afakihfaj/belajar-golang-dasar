package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}
	//endApp() // kalau panggil endApp() di sini tanpa defer, dia ga bakal jalan karena program keburu mati kena panic duluan
}

func main() {
	runApp(true)
}
