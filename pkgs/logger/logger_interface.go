package logger

type SharedLogger interface {
	Debug(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Info(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Error(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}
