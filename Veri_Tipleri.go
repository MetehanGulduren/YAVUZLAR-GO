package main

import "fmt"

func main() {
	var durum bool

	var metin1 string = "metehan"
	var metin2 string = "ahmet"

	durum = metin1 == metin2

	fmt.Println(durum)

	var dürüm bool

	var metin3 string = "metehan"
	var metin4 string = "!metehan"

	dürüm = metin3 != metin4

	fmt.Println(dürüm)
}
