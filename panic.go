package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}
	// message := recover()
	// fmt.Println("terjadi panic", message) kalo ditaruh disini,recovernya ga ke panggil,taruh di endapp mas

	//endApp() // kalau panggil endApp() di sini tanpa defer, dia ga bakal jalan karena program keburu mati kena panic duluan
}

func main() {
	runApp(true)
	fmt.Println("Afakih Fajduwani")
}
