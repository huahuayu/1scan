package logger

import (
	"log"
	"os"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var (
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
	infoLogger  = log.New(os.Stdout, "INFO:  ", log.LstdFlags)
	warnLogger  = log.New(os.Stdout, "WARN:  ", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)

	currentLevel = InfoLevel // Default level
)

// SetLevel sets the logging level
func SetLevel(level Level) {
	currentLevel = level
}

func Debug(format string, v ...interface{}) {
	if currentLevel <= DebugLevel {
		debugLogger.Printf(format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if currentLevel <= InfoLevel {
		infoLogger.Printf(format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	if currentLevel <= WarnLevel {
		warnLogger.Printf(format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if currentLevel <= ErrorLevel {
		errorLogger.Printf(format, v...)
	}
}
