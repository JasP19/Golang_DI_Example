package utilities

import "crypto/cipher"

// Encryption interface
type Encryption interface {
	EncryptData(data string) cipher.Block
}

func provideEncrypt() *Encrypt {
	en := new(Encrypt)
	return en
}
