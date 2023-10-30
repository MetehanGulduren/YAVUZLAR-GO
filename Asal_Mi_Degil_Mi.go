package main

import "fmt"

func main() {
	sayi := 0
	fmt.Println("bir sayi girin")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%1 == 0 {
			asalMi = false

		}

	}
	if asalMi == true {
		fmt.Println("asaldır")

	} else {
		fmt.Println("asal değildir")
	}
}
