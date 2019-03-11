//+build wireinject

package main

import (
	"github.com/google/wire"

	"com.softwarethree/IndividualProject/utilities"
)

// InitTestAccount injector
func InitTestAccount() utilities.Account {
	testAccount := utilities.ProvideTestAccount()
	return testAccount
}

// InitSummary injector
func InitSummary(messagePhrase string) utilities.AccountSummary {
	wire.Build(utilities.ProvideAccountSummary, utilities.ProvideAccount)

	return utilities.AccountSummary{}
}

var EncryptSet = wire.NewSet(
	utilities.ProvideEncrypt,
	wire.Bind((*utilities.Encryption)(nil), (*utilities.Encrypt)(nil)),
	utilities.ProvideEncryptor)

// InitEncryption injector
func InitEncryption() utilities.Encrypt {
	wire.Build(EncryptSet)

	return utilities.Encrypt{}
}
