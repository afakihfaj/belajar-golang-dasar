package main

import "fmt"

func main() {
	names := [...]string{"Afakih", "Fajduwani", "Nur", "Azka", "Diana", "Ayudia"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	angka := [...]int{1, 2, 3, 4, 5, 6}

	slice6 := angka[4:6]
	fmt.Println(slice6)

	days := [...]string{"Senin", "Selasa", "Rabu", "kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Afakih"
	newSlice[1] = "Afakih"
	// newSlice[2] = "Afakih" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Afakih")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Dewangga"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
