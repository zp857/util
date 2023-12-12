package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(plaintext []byte) string {
	hash := sha1.Sum(plaintext)
	return hex.EncodeToString(hash[:])
}

func VerifySha1(hash string, password string) bool {
	passwdHash := Sha1([]byte(password))
	if hash == passwdHash {
		return true
	} else {
		return false
	}
}
