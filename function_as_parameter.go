package main

import "fmt"

type Filter func(string) string // type alias buat fungsi: biar deklarasi parameter fungsi ga kepanjangan nulis func(string) string
// bikin cetakan/alias: biar bentuk func(string) string punya nama panggilan 'Filter' dan kodenya rapi

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)

}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Afakih", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
