package faker

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomString(length int) string {

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandomEmail() string {
	return GetRandomString(10) + "@email.com"
}

func GetRandomBool() bool {
	return rand.Intn(2) == 1
}
