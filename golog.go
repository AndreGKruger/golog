package golog

import (
	"fmt"
	"os"
	"time"
)

type Log interface {
	Debug(message string, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
	Configure(*Config) (bool, error)
	shouldLog(logtype int) bool
}

type log struct {
	config *Config
}

// Returns a new Log instance with default configuration
func New() Log {
	cnf := Config{}
	cnf.setupDefault()
	return &log{config: &cnf}
}

// Custom configuration for the logger
func (l *log) Configure(config *Config) (bool, error) {
	if valid, err := config.isValid(); !valid {
		return false, err
	}
	if config.LogFileName != "" {
		l.config.LogFileName = config.LogFileName
	}
	err := config.configureLogTo(config)
	if err != nil {
		return false, err
	}
	if config.LogEnvironment != "" {
		l.config.LogEnvironment = config.LogEnvironment
	}
	if os.Getenv("GOLOG_ENV") != "" {
		if os.Getenv("GOLOG_ENV") == DEVELOPMENT || os.Getenv("GOLOG_ENV") == STAGING || os.Getenv("GOLOG_ENV") == PRODUCTION {
			l.config.LogEnvironment = os.Getenv("GOLOG_ENV")
		}
	}
	if config.OutputFormat != 0 {
		l.config.OutputFormat = config.OutputFormat
	}
	return true, nil
}

func (l *log) Debug(message string, args ...any) {
	if l.shouldLog(DEBUG) {
		if args != nil {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [DEBUG]: %s  - [args]: %d\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message, args)
		} else {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [DEBUG]: %s\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message)
		}
	}
}

func (l *log) Info(message string, args ...any) {
	if l.shouldLog(INFO) {
		if args != nil {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [INFO]: %s  - [args]: %d\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message, args)
		} else {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [INFO]: %s\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message)
		}
	}
}

func (l *log) Warn(message string, args ...any) {
	if l.shouldLog(WARN) {
		if args != nil {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [WARN]: %s  - [args]: %d\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message, args)
		} else {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [WARN]: %s\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message)
		}
	}
}

func (l *log) Error(message string, args ...any) {
	if l.shouldLog(ERROR) {
		if args != nil {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [ERROR]: %s  - [args]: %d\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message, args)
		} else {
			fmt.Fprintf(l.config.writer, "[%s] - [ENV:%s] - [ERROR]: %s\n", time.Now().Format(time.RFC1123), l.config.LogEnvironment, message)
		}
	}
}

func (l *log) shouldLog(logtype int) bool {
	switch logtype {
	case DEBUG:
		return l.config.LogEnvironment == DEVELOPMENT
	case INFO:
		return l.config.LogEnvironment == DEVELOPMENT || l.config.LogEnvironment == STAGING
	case WARN:
		return l.config.LogEnvironment == DEVELOPMENT || l.config.LogEnvironment == STAGING || l.config.LogEnvironment == PRODUCTION
	case ERROR:
		return l.config.LogEnvironment == DEVELOPMENT || l.config.LogEnvironment == STAGING || l.config.LogEnvironment == PRODUCTION
	default:
		return false
	}
}
