package main

import (
	"fmt"
)

// driver code
func main() {
	// encryptor := utilities.Encrypt{}
	// cipher := encryptor.EncryptData("Hi")
	// fmt.Println(cipher)

	as := InitSummary()

	fmt.Println(as.GetSummary())
}
