package main

import "fmt"

func main() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("sayi tahmin et")
	fmt.Scanln(&tahminEdilenSayi)

	if tahminEdilenSayi < aklimdakiSayi {
		fmt.Println("daha büyük bir sayi gir")
	}
	if tahminEdilenSayi > aklimdakiSayi {
		fmt.Println("daha kucuk bir sayi gir")
	}
	if tahminEdilenSayi == aklimdakiSayi {
		fmt.Println("sayiyi tutturdun bravo")
	}

}
