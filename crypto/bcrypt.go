package crypto

/*

References:
https://pkg.go.dev/golang.org/x/crypto/bcrypt
https://www.educative.io/answers/how-to-use-bcrypt-to-hash-in-go

*/

import (
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(plaintext []byte) ([]byte, error) { // This function receives bytes and returns hashed bytes
	hash, err := bcrypt.GenerateFromPassword(plaintext, bcrypt.DefaultCost) // Generate hash
	if err != nil {
		return []byte(""), err
	}

	return hash, nil // Return hash w/o errors
}

func VerifyBcrypt(hash []byte, plaintext []byte) bool {
	hashCheck := bcrypt.CompareHashAndPassword(hash, plaintext)
	if hashCheck == nil {
		return true
	}
	return false
}
