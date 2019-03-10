// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"com.softwarethree/IndividualProject/utilities"
)

// Injectors from wire.go:

func InitSummary() utilities.AccountSummary {
	account := utilities.ProvideAccount()
	accountSummary := utilities.ProvideAccountSummary(account)
	return accountSummary
}
