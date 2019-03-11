package utilities

import (
	"crypto/cipher"
	"crypto/des"
)

// Encrypt service
type Encrypt struct {
	data string
	key  []byte
}

// ProvideEncrypt method
func ProvideEncrypt() *Encrypt {
	en := &Encrypt{data: "", key: []byte{}}
	return en
}

// ProvideEncryptor method
func ProvideEncryptor() Encrypt {
	return Encrypt{data: "", key: []byte{}}
}

// EncryptData Method
func (e *Encrypt) EncryptData(data string) cipher.Block {
	e.key = []byte{0xAC, 0xBC, 0xCC, 0xDC, 0xEC, 0xFC, 0xAD, 0xBD}
	blockDes, err := des.NewCipher(e.key)

	if err != nil {
		panic(err)
	}

	return blockDes
}
