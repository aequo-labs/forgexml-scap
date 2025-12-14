package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String returns the string representation of a log level
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// SysLogger defines the interface for system logging
type SysLogger interface {
	Close() error
	Log(level LogLevel, message string) error
}

// Logger represents a custom logger with file rotation and system logging
type Logger struct {
	writer     io.Writer
	level      LogLevel
	lumberjack *lumberjack.Logger
	sysLogger  SysLogger
}

// Config holds configuration for creating a new logger
type Config struct {
	LogDir      string
	LogFileName string
	Level       LogLevel
	UseSysLog   bool
	MaxSize     int  // Maximum size in megabytes before rotation
	MaxBackups  int  // Maximum number of old log files to retain
	MaxAge      int  // Maximum number of days to retain old log files
	Compress    bool // Whether to compress old log files
	ToStdout    bool // Whether to also output to stdout
}

// DefaultConfig returns a default logger configuration
func DefaultConfig() Config {
	return Config{
		LogDir:      "logs",
		LogFileName: "app.log",
		Level:       INFO,
		UseSysLog:   false,
		MaxSize:     50,
		MaxBackups:  7,
		MaxAge:      7,
		Compress:    true,
		ToStdout:    true, // Default to true for better development experience
	}
}

// log writes a log message with the specified level and prefix
func (l *Logger) log(level LogLevel, prefix string, format string, v ...interface{}) {
	if level >= l.level {
		timeStr := time.Now().Format("2006/01/02 15:04:05")
		format = strings.TrimRight(format, "\n")
		msg := fmt.Sprintf(format, v...)
		formattedMsg := fmt.Sprintf("%s %s%s", timeStr, prefix, msg)

		// Write to the multi-writer (file and stdout)
		fmt.Fprintln(l.writer, formattedMsg)

		// If system logger is available, write to it
		if l.sysLogger != nil {
			if err := l.sysLogger.Log(level, formattedMsg); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write to system log: %v\n", err)
			}
		}
	}
}

// Debug logs a message at DEBUG level with optional key-value pairs
func (l *Logger) Debug(msg string, keyvals ...interface{}) {
	if len(keyvals) > 0 {
		// Structured logging format
		var parts []string
		parts = append(parts, msg)
		for i := 0; i < len(keyvals); i += 2 {
			if i+1 < len(keyvals) {
				parts = append(parts, fmt.Sprintf("%v=%v", keyvals[i], keyvals[i+1]))
			}
		}
		l.log(DEBUG, "[DEBUG] ", "%s", strings.Join(parts, " "))
	} else {
		l.log(DEBUG, "[DEBUG] ", "%s", msg)
	}
}

// Debugf logs a formatted message at DEBUG level
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(DEBUG, "[DEBUG] ", format, v...)
}

// Error logs a message at ERROR level
func (l *Logger) Error(v ...interface{}) {
	l.log(ERROR, "[ERROR] ", "%s", fmt.Sprint(v...))
}

// Errorf logs a formatted message at ERROR level
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(ERROR, "[ERROR] ", format, v...)
}

// Info logs a message at INFO level with optional key-value pairs
func (l *Logger) Info(msg string, keyvals ...interface{}) {
	if len(keyvals) > 0 {
		// Structured logging format
		var parts []string
		parts = append(parts, msg)
		for i := 0; i < len(keyvals); i += 2 {
			if i+1 < len(keyvals) {
				parts = append(parts, fmt.Sprintf("%v=%v", keyvals[i], keyvals[i+1]))
			}
		}
		l.log(INFO, "[INFO] ", "%s", strings.Join(parts, " "))
	} else {
		l.log(INFO, "[INFO] ", "%s", msg)
	}
}

// Infof logs a formatted message at INFO level
func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(INFO, "[INFO] ", format, v...)
}

// Warn logs a message at WARN level
func (l *Logger) Warn(v ...interface{}) {
	l.log(WARN, "[WARN] ", "%s", fmt.Sprint(v...))
}

// Warnf logs a formatted message at WARN level
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log(WARN, "[WARN] ", format, v...)
}

// Fatal logs a message at FATAL level and exits the program
func (l *Logger) Fatal(v ...interface{}) {
	l.log(FATAL, "[FATAL] ", "%s", fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf logs a formatted message at FATAL level and exits the program
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(FATAL, "[FATAL] ", format, v...)
	os.Exit(1)
}

// Print logs a message without a level prefix
func (l *Logger) Print(v ...interface{}) {
	l.log(INFO, "", "%s", fmt.Sprint(v...))
}

// Printf logs a formatted message without a level prefix
func (l *Logger) Printf(format string, v ...interface{}) {
	l.log(INFO, "", format, v...)
}

// Println logs a message without a level prefix (with newline)
func (l *Logger) Println(v ...interface{}) {
	l.log(INFO, "", "%s", fmt.Sprintln(v...))
}

// Fatalln logs a message at FATAL level (with newline) and exits the program
func (l *Logger) Fatalln(v ...interface{}) {
	l.log(FATAL, "[FATAL] ", "%s", fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic logs a message at ERROR level and panics
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log(ERROR, "[PANIC] ", "%s", s)
	panic(s)
}

// Panicf logs a formatted message at ERROR level and panics
func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.log(ERROR, "[PANIC] ", "%s", s)
	panic(s)
}

// Panicln logs a message at ERROR level (with newline) and panics
func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.log(ERROR, "[PANIC] ", "%s", s)
	panic(s)
}

// Close closes the logger and all its resources
func (l *Logger) Close() {
	if l.lumberjack != nil {
		l.lumberjack.Close()
	}
	if l.sysLogger != nil {
		l.sysLogger.Close()
	}
}

// NewLogger creates a new logger with the specified configuration
func NewLogger(config Config) (*Logger, error) {
	if err := os.MkdirAll(config.LogDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	logFile := filepath.Join(config.LogDir, config.LogFileName)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	var sysLogger SysLogger
	var err error
	if config.UseSysLog {
		sysLogger, err = newSysLogger(config.LogFileName)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize system logger: %v", err)
		}
	}

	var logWriter io.Writer
	if config.ToStdout {
		logWriter = io.MultiWriter(os.Stdout, lumberjackLogger)
	} else {
		logWriter = lumberjackLogger
	}

	logger := &Logger{
		writer:     logWriter,
		level:      config.Level,
		lumberjack: lumberjackLogger,
		sysLogger:  sysLogger,
	}

	logger.Info("Logging system initialized with level:", config.Level)

	return logger, nil
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// GetLevel returns the current log level
func (l *Logger) GetLevel() LogLevel {
	return l.level
}