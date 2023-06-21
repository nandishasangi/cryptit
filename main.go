package main

import (
	"fmt"
	"github.com/nandishasangi/cryptit/encrypt"
	"github.com/nandishasangi/cryptit/decrypt"
)

func main(){
	fmt.Println(encrypt.Nimbus("Bijapur, Karnataka"))
	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("Bijapur, Karnataka")))
}