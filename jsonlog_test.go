package log

import (
	"bytes"
	golog "log"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// Test default configuration
	defaultLogger := defaultJSONLogger()
	if defaultLogger == nil {
		t.Error("Expected non-nil logger, got nil")
	}

	// Test custom configuration
	prefix := "test"
	verbosity := DEBUG
	cfg := &LoggerConfiguration{Prefix: prefix, Verbosity: verbosity}
	logger := newLogger(cfg)
	if logger == nil {
		t.Error("Expected non-nil logger, got nil")
	}
	if logger.prefix != prefix {
		t.Errorf("Expected prefix %q, got %q", prefix, logger.prefix)
	}
	if logger.verbosity != verbosity {
		t.Errorf("Expected verbosity %d, got %d", verbosity, logger.verbosity)
	}
}

func TestShutdown(t *testing.T) {
	logger := defaultJSONLogger()
	err := logger.Shutdown()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func TestSetVerbosity(t *testing.T) {
	logger := defaultJSONLogger()
	logger.SetVerbosity("debug")
	if logger.verbosity != DEBUG {
		t.Errorf("Expected verbosity %d, got %d", DEBUG, logger.verbosity)
	}
	logger.SetVerbosity("info")
	if logger.verbosity != INFO {
		t.Errorf("Expected verbosity %d, got %d", INFO, logger.verbosity)
	}
	logger.SetVerbosity("warn")
	if logger.verbosity != WARN {
		t.Errorf("Expected verbosity %d, got %d", WARN, logger.verbosity)
	}
	logger.SetVerbosity("err")
	if logger.verbosity != ERR {
		t.Errorf("Expected verbosity %d, got %d", ERR, logger.verbosity)
	}
	logger.SetVerbosity("fatal")
	if logger.verbosity != FATAL {
		t.Errorf("Expected verbosity %d, got %d", FATAL, logger.verbosity)
	}
	logger.SetVerbosity("invalid")
	if logger.verbosity != WARN {
		t.Errorf("Expected verbosity %d, got %d", WARN, logger.verbosity)
	}
}

func TestDebug(t *testing.T) {
	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	logger := &logger{out: golog.New(buf, "", 0), err: golog.New(bufErr, "", 0)}
	logger.SetVerbosity("debug")
	logger.Debug("test")
	if !strings.Contains(buf.String(), "debug") {
		t.Error("Expected log message to contain 'debug'")
	}
	if !strings.Contains(buf.String(), "test") {
		t.Error("Expected log message to contain 'test'")
	}
	if bufErr.String() != "" {
		t.Error("Expected empty log message in error output, got", buf.String())
	}
	buf.Reset()
	logger.verbosity = INFO
	logger.Debug("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
}

func TestInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	logger := &logger{out: golog.New(buf, "", 0), err: golog.New(bufErr, "", 0)}
	logger.verbosity = INFO
	logger.Info("test")
	if !strings.Contains(buf.String(), "info") {
		t.Error("Expected log message to contain 'info'")
	}
	if !strings.Contains(buf.String(), "test") {
		t.Error("Expected log message to contain 'test'")
	}
	buf.Reset()
	logger.Debug("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	if bufErr.String() != "" {
		t.Error("Expected empty log message in error output, got", buf.String())
	}
	logger.verbosity = WARN
	logger.Info("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
}

func TestWarn(t *testing.T) {
	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	logger := &logger{out: golog.New(buf, "", 0), err: golog.New(bufErr, "", 0)}
	logger.verbosity = WARN
	logger.Warn("test")
	if !strings.Contains(buf.String(), "warn") {
		t.Error("Expected log message to contain 'warn'")
	}
	if !strings.Contains(buf.String(), "test") {
		t.Error("Expected log message to contain 'test'")
	}
	buf.Reset()
	logger.Err("test")
	if !strings.Contains(bufErr.String(), "error") {
		t.Error("Expected log message to contain 'error'")
	}
	if !strings.Contains(bufErr.String(), "test") {
		t.Error("Expected log message to contain 'test'")
	}
	bufErr.Reset()
	logger.Debug("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.Info("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	if bufErr.String() != "" {
		t.Error("Expected empty log message in error output, got", bufErr.String())
	}
	logger.verbosity = ERR
	logger.Warn("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
}

func TestErr(t *testing.T) {
	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	logger := &logger{out: golog.New(buf, "", 0), err: golog.New(bufErr, "", 0)}
	logger.verbosity = ERR
	logger.Err("test")
	if !strings.Contains(bufErr.String(), "error") {
		t.Error("Expected log message to contain 'error'")
	}
	if !strings.Contains(bufErr.String(), "test") {
		t.Error("Expected log message to contain 'test'")
	}
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	bufErr.Reset()
	logger.Debug("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.Info("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.Warn("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.verbosity = FATAL
	logger.Err("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
}

func TestPanic(t *testing.T) {
	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	logger := &logger{out: golog.New(buf, "", 0), err: golog.New(bufErr, "", 0)}
	logger.verbosity = FATAL
	logger.Debug("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.Info("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.Warn("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	logger.Err("test")
	if buf.String() != "" {
		t.Error("Expected empty log message, got", buf.String())
	}
	buf.Reset()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic")
		}
		if !strings.Contains(bufErr.String(), "panic") {
			t.Error("Expected log message to contain 'panic'")
		}
		if !strings.Contains(bufErr.String(), "test") {
			t.Error("Expected log message to contain 'test'")
		}
	}()
	logger.Panic("test")
}

func TestFatal(t *testing.T) {
	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	logger := &logger{out: golog.New(buf, "", 0), err: golog.New(bufErr, "", 0)}
	logger.verbosity = FATAL
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic")
		}
		if !strings.Contains(bufErr.String(), "fatal") {
			t.Error("Expected log message to contain 'fatal'")
		}
		if !strings.Contains(bufErr.String(), "test") {
			t.Error("Expected log message to contain 'test'")
		}
	}()
	logger.Fatal("test")
}
