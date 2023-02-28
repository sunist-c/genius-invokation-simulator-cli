package log

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type logContext struct {
	prefix string
	format string
	args   []interface{}
}

type loggerImpl struct {
	started   bool
	logOut    map[string]io.Writer
	logBuffer chan logContext
	errOut    map[string]io.Writer
	errBuffer chan logContext
	exitChan  chan struct{}
}

func (impl *loggerImpl) formatString(prefix, format string, args ...interface{}) string {
	return fmt.Sprintf("%v[%v]: %v", prefix, time.Now().Format("2006.01.02 15:04:05"), fmt.Sprintf(format, args...))
}

func (impl *loggerImpl) formatSlice(slice ...interface{}) (format string, args []interface{}) {
	length := len(slice)
	formats := make([]string, length)
	for i := 0; i < length; i++ {
		formats[i] = "%v"
	}

	return strings.Join(formats, ", "), slice
}

func (impl *loggerImpl) serve() {
	if impl.started {
		return
	} else {
		impl.started = true
	}

	for {
		select {
		case ctx := <-impl.logBuffer:
			content := fmt.Sprintf("%v\n", impl.formatString(ctx.prefix, ctx.format, ctx.args...))
			for _, writer := range impl.logOut {
				if _, err := writer.Write([]byte(content)); err != nil {
					impl.Errorf("error writing log: %v", err)
				}
			}
		case ctx := <-impl.errBuffer:
			content := fmt.Sprintf("%v\n", impl.formatString(ctx.prefix, ctx.format, ctx.args...))
			for _, writer := range impl.errOut {
				if _, err := writer.Write([]byte(content)); err != nil {
					impl.Panicf("error writing log: %v", err)
				}
			}
		case <-impl.exitChan:
			impl.started = false
			return
		}
	}
}

func (impl *loggerImpl) Logf(format string, args ...interface{}) {
	impl.logBuffer <- logContext{
		prefix: "[Debug]",
		format: format,
		args:   args,
	}
}

func (impl *loggerImpl) Log(content ...interface{}) {
	format, args := impl.formatSlice(content...)
	impl.logBuffer <- logContext{
		prefix: "[Debug]",
		format: format,
		args:   args,
	}
}

func (impl *loggerImpl) Errorf(format string, args ...interface{}) {
	impl.errBuffer <- logContext{
		prefix: "[Error]",
		format: format,
		args:   args,
	}
}

func (impl *loggerImpl) Error(error ...interface{}) {
	format, args := impl.formatSlice(error...)
	impl.errBuffer <- logContext{
		prefix: "[Error]",
		format: format,
		args:   args,
	}
}

func (impl *loggerImpl) Panicf(format string, args ...interface{}) {
	panic(impl.formatString("[Fatal]", format, args))
}

func (impl *loggerImpl) Panic(fatal ...interface{}) {
	format, args := impl.formatSlice(fatal...)
	panic(impl.formatString("[Fatal]", format, args))
}

type LoggerOptions func(option *loggerImpl)

func WithLoggerLogWriter(name string, writer io.Writer) LoggerOptions {
	return func(option *loggerImpl) {
		option.logOut[name] = writer
	}
}

func WithLoggerErrorWriter(name string, writer io.Writer) LoggerOptions {
	return func(option *loggerImpl) {
		option.errOut[name] = writer
	}
}

func NewLoggerWithOpts(options ...LoggerOptions) Logger {
	impl := &loggerImpl{
		started:   false,
		logOut:    map[string]io.Writer{},
		logBuffer: make(chan logContext, 8),
		errOut:    map[string]io.Writer{},
		errBuffer: make(chan logContext, 4),
		exitChan:  make(chan struct{}),
	}

	for _, option := range options {
		option(impl)
	}

	go impl.serve()

	return impl
}
