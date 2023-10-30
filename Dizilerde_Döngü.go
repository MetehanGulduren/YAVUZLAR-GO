package main

import "fmt"

func main() {
	var sehirler [5]string
	sehirler[0] = "antalya"
	sehirler[1] = "zonguldak"
	sehirler[2] = "bartın"
	sehirler[3] = "erzurum"
	sehirler[4] = "ankara"
	//fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
