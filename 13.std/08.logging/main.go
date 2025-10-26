package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

// LogLevel is a typed enum for allowed log levels.
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (l LogLevel) color() string {
	switch l {
	case INFO:
		return colorGreen
	case WARN:
		return colorYellow
	case ERROR:
		return colorRed
	default:
		return colorReset
	}
}

// LevelLogger is a small wrapper around the standard log.Logger that implements level filtering
// and writes to both a terminal logger (colored) and a file logger (plain, timestamped).
type LevelLogger struct {
	min   LogLevel
	term  *log.Logger
	file  *log.Logger
	termW io.Writer
	fileW io.Writer
}

// NewLevelLogger constructs a LevelLogger.
// - min: minimum level to emit (inclusive).
// - termOut: terminal writer (e.g., os.Stdout).
// - fileOut: file writer (opened *os.File).
func NewLevelLogger(min LogLevel, termOut io.Writer, fileOut io.Writer) *LevelLogger {
	termLogger := log.New(termOut, "", log.LstdFlags|log.Lshortfile)
	fileLogger := log.New(fileOut, "", log.LstdFlags|log.Lshortfile)
	return &LevelLogger{
		min:   min,
		term:  termLogger,
		file:  fileLogger,
		termW: termOut,
		fileW: fileOut,
	}
}

// shouldEmit checks whether a message of level should be emitted.
func (l *LevelLogger) shouldEmit(level LogLevel) bool {
	return level >= l.min
}

// Log writes a message for the given level to terminal and file according to the min level.
// Terminal output is colored (timestamp handled by term logger flags if set).
// File output is plain and timestamped by file logger flags.
func (l *LevelLogger) Log(level LogLevel, msg string) {
	if !l.shouldEmit(level) {
		return
	}

	// Terminal
	colored := fmt.Sprintf("%s[%s] %s%s", level.color(), level.String(), msg, colorReset)
	l.term.Println(colored)

	// File
	plain := fmt.Sprintf("[%s] %s", level.String(), msg)
	l.file.Println(plain)
}

func (l *LevelLogger) Info(msg string)  { l.Log(INFO, msg) }
func (l *LevelLogger) Warn(msg string)  { l.Log(WARN, msg) }
func (l *LevelLogger) Error(msg string) { l.Log(ERROR, msg) }

// createLogFile ensures the directory exists and opens the log file for append.
func createLogFile(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	// open log file
	logPath := "main.log"
	f, err := createLogFile(logPath)
	if err != nil {
		log.Fatalf("failed to create/open log file %s: %v", logPath, err)
	}
	defer f.Close()

	//logger := NewLevelLogger(INFO, os.Stdout, f)
	logger := NewLevelLogger(ERROR, os.Stdout, f)
	logger.Info("Application healthy")
	logger.Warn("Cache miss for key 'user:1234'")
	logger.Error("Failed to connect to database")
}
