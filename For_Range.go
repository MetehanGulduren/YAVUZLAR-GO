package main

import "fmt"

func main() {
	sehirler := []string{"Ankara", "istanbul", "izmir"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}

}
