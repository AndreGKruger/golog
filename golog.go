package golog

import (
	"log"
	"os"
)

type Log interface {
	Debug(message string, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
	Configure(*Config) (bool, error)
	shouldLog(logtype int) bool
}

type golog struct {
	config *Config
	log    *log.Logger
}

// Returns a new Log instance with default configuration
func New() Log {
	cnf := Config{}
	cnf.setupDefault()
	return &golog{
		config: &cnf,
		log:    log.New(cnf.writer, "[ ", log.LstdFlags),
	}
}

// Custom configuration for the logger
func (l *golog) Configure(config *Config) (bool, error) {
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
	l.log.SetOutput(config.writer)
	if config.LogEnvironment != "" {
		l.config.LogEnvironment = config.LogEnvironment
	}
	if os.Getenv("GOLOG_ENV") != "" {
		if os.Getenv("GOLOG_ENV") == CONFIG_ENV_DEVELOPMENT || os.Getenv("GOLOG_ENV") == CONFIG_ENV_STAGING || os.Getenv("GOLOG_ENV") == CONFIG_ENV_PRODUCTION {
			l.config.LogEnvironment = os.Getenv("GOLOG_ENV")
		}
	}
	if config.OutputFormat != 0 {
		l.config.OutputFormat = config.OutputFormat
	}
	return true, nil
}

func (l *golog) Debug(message string, args ...any) {
	if l.shouldLog(CONFIG_LOG_LEVEL_DEBUG) {
		if args != nil {
			l.log.Printf("] - [ENV:%s] - [DEBUG]: %s  - [ARGS]: %d\n", l.config.LogEnvironment, message, args)
		} else {
			l.log.Printf("] - [ENV:%s] - [DEBUG]: %s\n", l.config.LogEnvironment, message)
		}
	}
}

func (l *golog) Info(message string, args ...any) {
	if l.shouldLog(CONFIG_LOG_LEVEL_INFO) {
		if args != nil {
			l.log.Printf("] - [ENV:%s] - [INFO]: %s  - [ARGS]: %d\n", l.config.LogEnvironment, message, args)
		} else {
			l.log.Printf("] - [ENV:%s] - [INFO]: %s\n", l.config.LogEnvironment, message)
		}
	}
}

func (l *golog) Warn(message string, args ...any) {
	if l.shouldLog(CONFIG_LOG_LEVEL_WARN) {
		if args != nil {
			l.log.Printf("] - [ENV:%s] - [WARN]: %s  - [ARGS]: %d\n", l.config.LogEnvironment, message, args)
		} else {
			l.log.Printf("] - [ENV:%s] - [WARN]: %s\n", l.config.LogEnvironment, message)
		}
	}
}

func (l *golog) Error(message string, args ...any) {
	if l.shouldLog(CONFIG_LOG_LEVEL_ERROR) {
		if args != nil {
			l.log.Printf("] - [ENV:%s] - [ERROR]: %s  - [ARGS]: %d\n", l.config.LogEnvironment, message, args)
		} else {
			l.log.Printf("] - [ENV:%s] - [ERROR]: %s\n", l.config.LogEnvironment, message)
		}
	}
}

func (l *golog) shouldLog(logtype int) bool {
	switch logtype {
	case CONFIG_LOG_LEVEL_DEBUG:
		return l.config.LogEnvironment == CONFIG_ENV_DEVELOPMENT
	case CONFIG_LOG_LEVEL_INFO:
		return l.config.LogEnvironment == CONFIG_ENV_DEVELOPMENT || l.config.LogEnvironment == CONFIG_ENV_STAGING
	case CONFIG_LOG_LEVEL_WARN:
		return l.config.LogEnvironment == CONFIG_ENV_DEVELOPMENT || l.config.LogEnvironment == CONFIG_ENV_STAGING || l.config.LogEnvironment == CONFIG_ENV_PRODUCTION
	case CONFIG_LOG_LEVEL_ERROR:
		return l.config.LogEnvironment == CONFIG_ENV_DEVELOPMENT || l.config.LogEnvironment == CONFIG_ENV_STAGING || l.config.LogEnvironment == CONFIG_ENV_PRODUCTION
	default:
		return false
	}
}
