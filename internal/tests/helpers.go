package tests

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/hay-kot/git-web-template/pkgs/logger"
)

func GetConsoleLogger(t *testing.T) logger.SharedLogger {
	return logger.NewStandardLogger(log.New(os.Stdout, "", 0))
}

func GetRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandomEmail() string {
	return GetRandomString(10) + "@email.com"
}
