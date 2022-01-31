package mocks

import (
	"os"

	"github.com/hay-kot/git-web-template/pkgs/logger"
)

func GetStructLogger() *logger.Logger {
	return logger.New(os.Stdout, logger.LevelDebug)
}
