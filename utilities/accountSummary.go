package utilities

import "strconv"

// AccountSummary struct
type AccountSummary struct {
	account Account
	message string
}

// ProvideAccountSummary method
func ProvideAccountSummary(a Account) AccountSummary {
	return AccountSummary{account: a, message: "Welcome"}
}

// GetSummary output
func (as AccountSummary) GetSummary() string {
	return (as.message + " " + as.account.name + ". Your balance is R" + strconv.FormatFloat(as.account.balance, 'f', -1, 64))
}
