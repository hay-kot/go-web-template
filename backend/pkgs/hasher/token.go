package hasher

import "golang.org/x/crypto/bcrypt"

func NewToken() string {
	s, _ := GenerateRandomString(32)
	return s
}

func HashToken(token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckTokenHash(token, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err == nil
}
