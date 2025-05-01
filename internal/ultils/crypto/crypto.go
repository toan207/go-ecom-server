package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
