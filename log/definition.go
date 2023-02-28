package log

type Logger interface {
	Logf(format string, args ...interface{})
	Log(content ...interface{})
	Errorf(format string, args ...interface{})
	Error(error ...interface{})
	Panicf(format string, args ...interface{})
	Panic(fatal ...interface{})
}
