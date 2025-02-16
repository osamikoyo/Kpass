package crypto

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)


func Encrypt(data string, key []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil{
		return "", fmt.Errorf("cant get cipher blocks: %w", err)
	}

	out := make([]byte, len(data))

	c.Encrypt(out, []byte(data))

	return hex.EncodeToString(out), nil
}