package main

import "fmt"

func main() {
	var hesap float64 = 1000
	var cekilmeIstenen float64 = 900

	if cekilmeIstenen > hesap {
		fmt.Print("para yetersiz")
	}
	fmt.Println("bitti")
	if cekilmeIstenen <= hesap {
		fmt.Println("paranız hazırlanıyor")
		hesap = hesap - cekilmeIstenen
	}
	fmt.Println("bitti. Hesaptaki para :" + fmt.Sprintf("%f", hesap))
}
