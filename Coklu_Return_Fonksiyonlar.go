package main

import "fmt"

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)

	fmt.Println("toplama:", toplam)
	fmt.Println("çıkarma:", cikarim)
	fmt.Println("çarpma:", carpim)
	fmt.Println("bölme:", bolum)

	return toplam, cikarim, carpim, bolum
}
