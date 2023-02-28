package log

import (
	"os"
	"path"
)

var (
	nilLogger       Logger = nil
	defaultImpl     Logger
	outputDirectory string
	consoleLog      bool
)

func SetOutputDirectory(directory string) {
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if mdErr := os.MkdirAll(directory, 0755); mdErr != nil {
				panic(mdErr)
			}
		} else {
			panic(err)
		}
	}

	outputDirectory = directory
}

func SetConsoleLog(enable bool) {
	consoleLog = enable
}

func init() {
	SetOutputDirectory("./logs")
	SetConsoleLog(false)
}

func Default() Logger {
	if defaultImpl == nilLogger {
		logFile, err := os.OpenFile(path.Join(outputDirectory, "stdout.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		errFile, err := os.OpenFile(path.Join(outputDirectory, "stderr.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}

		if consoleLog {
			defaultImpl = NewLoggerWithOpts(
				WithLoggerLogWriter("console", os.Stdout),
				WithLoggerLogWriter("file", logFile),
				WithLoggerErrorWriter("error", os.Stderr),
				WithLoggerErrorWriter("logfile", logFile),
				WithLoggerErrorWriter("file", errFile),
			)
		} else {
			defaultImpl = NewLoggerWithOpts(
				WithLoggerLogWriter("file", logFile),
				WithLoggerErrorWriter("logfile", logFile),
				WithLoggerErrorWriter("file", errFile),
			)
		}

		defaultImpl.Logf("initialized logger")
	}

	return defaultImpl
}
