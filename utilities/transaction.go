package utilities

// Transaction struct
type Transaction struct {
	dateTime    int64
	amount      float64
	toAccount   Account
	fromAccount Account
}

// import "time"
// seconds since epoch for dateTime
// time.Now().Unix()
