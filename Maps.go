package main

import "fmt"

func main() {
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("eleman sayısı:", len(sozluk))
	delete(sozluk, "book")
	fmt.Println("eleman sayısı:", len(sozluk))

	deger, VarMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu:", VarMi)

}
