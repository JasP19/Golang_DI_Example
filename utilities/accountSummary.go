package utilities

import "strconv"

// AccountSummary struct
type AccountSummary struct {
	account Account
	message string
}

// ProvideAccountSummary method
func ProvideAccountSummary(a Account, messagePhrase string) AccountSummary {
	return AccountSummary{account: a, message: messagePhrase}
}

// GetSummary output
func (as AccountSummary) PrintSummary() string {
	return (as.message + " " + as.account.name + ". Your balance is R" + strconv.FormatFloat(as.account.balance, 'f', -1, 64))
}
