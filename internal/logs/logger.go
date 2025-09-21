// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// LogLevel defines log levels
type LogLevel int

const (
	// DebugLevel debug level log
	DebugLevel LogLevel = iota
	// InfoLevel information level log
	InfoLevel
	// WarnLevel warning level log
	WarnLevel
	// ErrorLevel error level log
	ErrorLevel
	// FatalLevel fatal level log, program will exit after logging
	FatalLevel
)

var (
	// Log level string mapping
	levelStrings = map[LogLevel]string{
		DebugLevel: "DEBUG",
		InfoLevel:  "INFO",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		FatalLevel: "FATAL",
	}

	// Current log level
	currentLevel = InfoLevel

	// Log output destination
	output io.Writer = os.Stdout

	// Log formatter
	logger *log.Logger

	// Mutex for concurrency safety
	mutex sync.Mutex
)

// LoggerConfig defines log configuration
type LoggerConfig struct {
	Level  string // Log level
	Output string // Log output file path, defaults to stdout
	Prefix string // Log prefix
}

// InitLogger initializes the logging system
func InitLogger(config *LoggerConfig) {
	mutex.Lock()
	defer mutex.Unlock()

	// Set log level
	if config != nil && config.Level != "" {
		setLogLevel(config.Level)
	}

	// Set log output
	var logOutput io.Writer
	if config != nil && config.Output != "" {
		// Ensure log directory exists
		logDir := filepath.Dir(config.Output)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create log directory: %v\n", err)
			logOutput = os.Stdout
		} else {
			// Open log file
			file, err := os.OpenFile(config.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
				logOutput = os.Stdout
			} else {
				// Output to both file and stdout
				logOutput = io.MultiWriter(os.Stdout, file)
			}
		}
	} else {
		logOutput = os.Stdout
	}

	output = logOutput

	// Set log prefix
	prefix := ""
	if config != nil && config.Prefix != "" {
		prefix = config.Prefix + " "
	}

	// Create logger
	logger = log.New(output, prefix, 0)
}

// setLogLevel sets the log level
func setLogLevel(levelStr string) {
	levelStr = strings.ToUpper(levelStr)
	for level, str := range levelStrings {
		if str == levelStr {
			currentLevel = level
			return
		}
	}
	// If the specified level is invalid, use Info level by default
	Info("Invalid log level: %s, using default level: INFO", levelStr)
}

// shouldLog checks if logs of the specified level should be recorded
func shouldLog(level LogLevel) bool {
	return level >= currentLevel
}

// formatMessage formats log messages
func formatMessage(level LogLevel, format string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	levelStr := levelStrings[level]

	message := format
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}

	return fmt.Sprintf("%s [%s] %s\n", timestamp, levelStr, message)
}

// Debug logs debug level messages
func Debug(format string, args ...interface{}) {
	if shouldLog(DebugLevel) {
		mutex.Lock()
		defer mutex.Unlock()
		if logger != nil {
			logger.Print(formatMessage(DebugLevel, format, args...))
		} else {
			fmt.Fprint(output, formatMessage(DebugLevel, format, args...))
		}
	}
}

// Info logs information level messages
func Info(format string, args ...interface{}) {
	if shouldLog(InfoLevel) {
		mutex.Lock()
		defer mutex.Unlock()
		if logger != nil {
			logger.Print(formatMessage(InfoLevel, format, args...))
		} else {
			fmt.Fprint(output, formatMessage(InfoLevel, format, args...))
		}
	}
}

// Warn logs warning level messages
func Warn(format string, args ...interface{}) {
	if shouldLog(WarnLevel) {
		mutex.Lock()
		defer mutex.Unlock()
		if logger != nil {
			logger.Print(formatMessage(WarnLevel, format, args...))
		} else {
			fmt.Fprint(output, formatMessage(WarnLevel, format, args...))
		}
	}
}

// Error logs error level messages
func Error(format string, args ...interface{}) {
	if shouldLog(ErrorLevel) {
		mutex.Lock()
		defer mutex.Unlock()
		if logger != nil {
			logger.Print(formatMessage(ErrorLevel, format, args...))
		} else {
			fmt.Fprint(output, formatMessage(ErrorLevel, format, args...))
		}
	}
}

// Fatal logs fatal level messages and exits the program
func Fatal(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	message := formatMessage(FatalLevel, format, args...)
	if logger != nil {
		logger.Print(message)
	} else {
		fmt.Fprint(output, message)
	}

	os.Exit(1)
}

// ParseLogLevelFlag parses command line log level argument
func ParseLogLevelFlag(flag *string) LogLevel {
	levelStr := strings.ToUpper(*flag)
	for level, str := range levelStrings {
		if str == levelStr {
			return level
		}
	}
	// Return Info level by default
	return InfoLevel
}