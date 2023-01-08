package log

import (
	gofmt "fmt"
	golog "log"
	"os"
	"strconv"
	"time"
)

// logger is a concrete implementation of the Logger interface that logs to stdout and stderr.
type logger struct {
	out       *golog.Logger
	err       *golog.Logger
	verbosity int
	prefix    string
}

// NewLogger returns a new logger instance with the given configuration.
func NewLogger(cfg *LoggerConfiguration) Logger {
	if cfg == nil {
		return defaultJSONLogger()
	}
	return newLogger(cfg)
}

// defaultJSONLogger returns a new logger instance with default configuration (verbosity level set to INFO).
func defaultJSONLogger() *logger {
	return newLogger(&LoggerConfiguration{Verbosity: INFO})
}

func newLogger(cfg *LoggerConfiguration) *logger {
	return &logger{out: golog.New(os.Stdout, "", 0), err: golog.New(os.Stderr, "", 0), verbosity: cfg.Verbosity, prefix: cfg.Prefix}
}

// Shutdown closes any resources used by the logger.
func (l *logger) Shutdown() error {
	return nil
}

// SetVerbosity sets the verbosity level of the logger.
func (l *logger) SetVerbosity(Verbosity string) {
	l.verbosity = GetVerbosityFromString(Verbosity)
}

// Debug logs a message at the DEBUG level.
func (l *logger) Debug(fmt string, v ...interface{}) {
	if l.verbosity <= DEBUG {
		l.out.Println(l.format("debug", fmt, v...))
	}
}

// INFO logs a message at the INFO level.
func (l *logger) Info(fmt string, v ...interface{}) {
	if l.verbosity <= INFO {
		l.out.Println(l.format("info", fmt, v...))
	}
}

// Warn logs a message at the WARN level.
func (l *logger) Warn(fmt string, v ...interface{}) {
	if l.verbosity <= WARN {
		l.out.Println(l.format("warn", fmt, v...))
	}
}

// Err logs a message at the ERR level.
func (l *logger) Err(fmt string, v ...interface{}) {
	if l.verbosity <= ERR {
		l.err.Println(l.format("error", fmt, v...))
	}
}

// Panic logs a message at the Panic level and then panics.
func (l *logger) Panic(fmt string, v ...interface{}) {
	l.err.Panicln(l.format("panic", fmt, v...))
}

// Fatal logs a message at the Fatal level and then calls os.Exit(1).
func (l *logger) Fatal(fmt string, v ...interface{}) {
	l.err.Fatalln(l.format("fatal", fmt, v...))
}

// format formats the log message with the given level and arguments.
func (l *logger) format(level string, fmt string, v ...interface{}) string {
	return `{"time": "` + time.Now().Format(time.RFC3339Nano) + `", "level": "` + level + `", "message": ` + strconv.Quote(gofmt.Sprintf("%s"+fmt, append([]interface{}{l.prefix}, v...)...)) + `}`
}
