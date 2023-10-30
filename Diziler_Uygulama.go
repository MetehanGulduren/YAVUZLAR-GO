package main

import "fmt"

func main() {
	sayilar := [5]int{20, 10, 5, 4, 23}

	enBuyuk := sayilar[0]
	//lenght(len) = uzunluk
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}
	fmt.Println("En Büyük :", enBuyuk)
}
