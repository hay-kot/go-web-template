package hasher

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"

	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	Raw  string
	Hash []byte
}

func GenerateToken() Token {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	plainText := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := HashToken(plainText)

	return Token{
		Raw:  plainText,
		Hash: hash,
	}
}

func HashToken(plainTextToken string) []byte {
	hash := sha256.Sum256([]byte(plainTextToken))
	return hash[:]
}

func CheckTokenHash(token, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err == nil
}
