package utilities

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
