package utilities

import "strconv"

// Account struct
type Account struct {
	name       string
	accountNum string
	balance    float64
}

// ProvideAccount method
func ProvideAccount() Account {
	return Account{name: "Jason", accountNum: "1234", balance: 123.45}
}

// ProvideTestAccount method
func ProvideTestAccount() Account {
	return Account{name: "Jill", accountNum: "7890", balance: 789.00}
}

// PrintDetails method
func (a Account) PrintDetails() string {
	return "Name: " + a.name + "\nAccount Number: " + a.accountNum + "\nBalance: " + strconv.FormatFloat(a.balance, 'f', -1, 64)
}
