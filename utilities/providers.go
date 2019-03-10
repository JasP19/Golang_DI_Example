package utilities

// ProvideAccount method
func ProvideAccount() Account {
	return Account{name: "Jason", accountNum: "1234", balance: 123.45}
}

// ProvideAccountSummary method
func ProvideAccountSummary(a Account) AccountSummary {
	return AccountSummary{account: a, message: "Welcome"}
}

// ProvideEncrypt method
func ProvideEncrypt() *Encrypt {
	en := new(Encrypt)
	return en
}
