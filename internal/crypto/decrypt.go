package crypto

import (
	"crypto/aes"
	"encoding/hex"
)

func Decrypt(data string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(data)
	if err != nil{
		return "", nil
	}

	c,err := aes.NewCipher(key)
	if err != nil{
		return "", nil
	}

	out := make([]byte, len(data))

	c.Decrypt(out, ciphertext)

	return string(out[:]), nil
}