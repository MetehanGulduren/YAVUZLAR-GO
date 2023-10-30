package main

import "fmt"

func main() {
	var hesap float64 = 900
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("para yetersiz")

	} else if cekilmekIstenen == hesap {
		fmt.Println("para hazırlanıyor")
	}

}
