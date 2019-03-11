package main

import (
	"fmt"
)

// driver code
func main() {

	// standard dependency injection
	testAccount := InitTestAccount()
	fmt.Println(testAccount.PrintDetails())

	fmt.Println()

	/* building dependency tree (account then summary), changing injector signature with message, and
	struct provider in account attribute of account summary */
	accountSummary := InitSummary("Hello")
	fmt.Println(accountSummary.PrintSummary())

	fmt.Println()

	//transaction setup

	// encryption
	en := InitEncryption()
	fmt.Println(en.EncryptData("Hello"))
}
