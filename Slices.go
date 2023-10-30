package main

import (
	"fmt"
)

func main() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "metehan"
	isimler[1] = "siber_vatan"
	isimler[2] = "yavuzlar"
	isimler = append(isimler, "zonguldak")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
