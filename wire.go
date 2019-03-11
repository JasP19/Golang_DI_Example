//+build wireinject

package main

import (
	"github.com/google/wire"

	"com.softwarethree/IndividualProject/utilities"
)

// InitSummary injector
func InitSummary() utilities.AccountSummary {
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
