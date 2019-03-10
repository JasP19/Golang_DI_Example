package utilities

import "strconv"

// AccountSummary struct
type AccountSummary struct {
	account Account
	message string
}

// GetSummary output
func (as AccountSummary) GetSummary() string {
	return (as.message + " " + as.account.name + ". Your balance is " + strconv.FormatFloat(as.account.balance, 'f', -1, 64))
}
