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

// do binding and stuff
// func InitEncryption() {
// 	wire.Build()
// }
