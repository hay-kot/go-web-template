package logger

import (
	"log"
	"time"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	DebugLevel = "[DEBUG]"
	InfoLevel  = "[INFO]"
	WarnLevel  = "[WARN]"
	ErrorLevel = "[ERROR]"
	FatalLevel = "[FATAL]"
)

type StandardLogger struct {
	logger *log.Logger
}

func NewStandardLogger(lgr *log.Logger) *StandardLogger {
	return &StandardLogger{
		logger: lgr,
	}
}

func (s StandardLogger) prefix() string {
	return time.Now().UTC().Format("2006/01/02 15:04:05") + " "
}

func (s StandardLogger) level(level string) string {
	var color string

	switch level {
	case DebugLevel:
		color = Gray
	case InfoLevel:
		color = Blue
	case WarnLevel:
		color = Yellow
	case ErrorLevel:
		color = Red
	case FatalLevel:
		color = Purple
	default:
		color = White
	}

	return s.prefix() + color + level + Reset + " "
}

func (s StandardLogger) Debug(format string, args ...interface{}) {
	s.logger.Printf(s.level(DebugLevel)+format, args...)
}

func (s StandardLogger) Debugf(format string, args ...interface{}) {
	s.logger.Printf(s.level(DebugLevel)+format, args...)
}

func (s StandardLogger) Info(format string, args ...interface{}) {
	s.logger.Printf(s.level(InfoLevel)+format, args...)
}

func (s StandardLogger) Infof(format string, args ...interface{}) {
	s.logger.Printf(s.level(InfoLevel)+format, args...)
}

func (s StandardLogger) Error(format string, args ...interface{}) {
	s.logger.Printf(s.level(ErrorLevel)+format, args...)
}

func (s StandardLogger) Errorf(format string, args ...interface{}) {
	s.logger.Printf(s.level(ErrorLevel)+format, args...)
}

func (s StandardLogger) Fatal(format string, args ...interface{}) {
	s.logger.Printf(s.level(FatalLevel)+format, args...)
}

func (s StandardLogger) Fatalf(format string, args ...interface{}) {
	s.logger.Printf(s.level(FatalLevel)+format, args...)
}
