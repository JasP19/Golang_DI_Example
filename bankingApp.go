package main

import (
	"fmt"
)

// driver code
func main() {
	// encryptor := utilities.Encrypt{}
	// cipher := encryptor.EncryptData("Hi")
	// fmt.Println(cipher)

	as := InitSummary("Hello")
	en := InitEncryption()

	fmt.Println(as.GetSummary())
	fmt.Println(en.EncryptData("Hello"))
}
