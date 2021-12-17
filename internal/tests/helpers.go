package tests

import (
	"log"
	"os"
	"testing"

	"github.com/hay-kot/git-web-template/pkgs/logger"
)

func GetConsoleLogger(t *testing.T) logger.SharedLogger {
	return logger.NewStandardLogger(log.New(os.Stdout, "", 0))
}
