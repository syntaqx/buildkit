package buildkit

import (
	"fmt"
	"log"
)

// Logger defines a standard interface for logging.
type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	Panic(msg string, keysAndValues ...interface{})
}

type defaultLogger struct{}

func (l *defaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.Println("[DEBUG]:", fmt.Sprint(msg, keysAndValues))
}

func (l *defaultLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Println("[INFO]:", fmt.Sprint(msg, keysAndValues))
}

func (l *defaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	log.Println("[WARN]:", fmt.Sprint(msg, keysAndValues))
}

func (l *defaultLogger) Error(msg string, keysAndValues ...interface{}) {
	log.Println("[ERROR]:", fmt.Sprint(msg, keysAndValues))
}

func (l *defaultLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatalln("[FATAL]:", fmt.Sprint(msg, keysAndValues))
}

func (l *defaultLogger) Panic(msg string, keysAndValues ...interface{}) {
	log.Panicln("[PANIC]:", fmt.Sprint(msg, keysAndValues))
}

// NewDefaultLogger creates a new default logger that uses the Go's standard log package.
func NewDefaultLogger() Logger {
	return &defaultLogger{}
}
