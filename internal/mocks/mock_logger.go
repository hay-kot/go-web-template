package mocks

import (
	"log"
	"os"

	"github.com/hay-kot/git-web-template/pkgs/logger"
)

func GetConsoleLogger() logger.SharedLogger {
	return logger.NewStandardLogger(log.New(os.Stdout, "", 0))
}
