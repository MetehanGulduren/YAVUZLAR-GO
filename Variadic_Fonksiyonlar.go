package main

import "fmt"

func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}
	return toplam
}

func main() {
	sonuc := ToplaVariadic(1, 4, 6, 3, 10)
	fmt.Println(sonuc)
}
