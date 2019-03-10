package main

import (
	"fmt"
	"com.softwarethree/IndividualProject/utilities"
)

func main() {
	encryptor := encryption.Encrypt{}
	cipher := encryptor.EncryptData("Hi")
	fmt.Println(cipher)
}
