package utilities

import "crypto/cipher"

// Encryption interface
type Encryption interface {
	EncryptData(data string) cipher.Block
}
