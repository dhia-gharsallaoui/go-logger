package log

import "strings"

const (
	DEBUG = iota
	INFO
	WARN
	ERR
	FATAL
)

// LoggerConfiguration represents the configuration for a logger instance.
type LoggerConfiguration struct {
	// Prefix is a string that is prepended to each log message.
	Prefix string
	// Verbosity is an integer representing the level of verbosity for the logger.
	// It should be one of the constants defined above (DEBUG, INFO, WARN, ERR, or FATAL).
	Verbosity int
}

// Logger is an interface that defines the methods that a logger should implement.
type Logger interface {
	// Debug logs a message at the Debug level.
	Debug(string, ...interface{})
	// Info logs a message at the Info level.
	Info(string, ...interface{})
	// Warn logs a message at the Warn level.
	Warn(string, ...interface{})
	// Err logs a message at the Err level.
	Err(string, ...interface{})
	// Panic logs a message at the Panic level and then panics.
	Panic(string, ...interface{})
	// Fatal logs a message at the Fatal level and then calls os.Exit(1).
	Fatal(string, ...interface{})
	// SetVerbosity sets the verbosity level of the logger.
	SetVerbosity(string)
	// Shutdown closes any resources used by the logger.
	Shutdown() error
}

// GetVerbosityFromString converts a string representation of a verbosity level
// to the corresponding integer constant.
// If the input string is not recognized, the function returns Warn.
func GetVerbosityFromString(verbosity string) int {
	switch strings.ToLower(verbosity) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "err":
		return ERR
	case "fatal":
		return FATAL
	default:
		return WARN
	}
}
