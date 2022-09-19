package golog

import (
	"fmt"
	"io"
	"os"
)

// Supported GOLOG_ENV values
const (
	DEVELOPMENT = "development"
	STAGING     = "staging"
	PRODUCTION  = "production"
)

const (
	// OutputFormat constants
	OUTPUT_FORMAT_JSON = iota + 1
	OUTPUT_FORMAT_TEXT
	// LogTo constants
	CONSOLE
	FILE
	DEBUG
	// Supported Logtypes
	INFO
	WARN
	ERROR
)

type Config struct {
	LogTo          uint
	OutputFormat   uint
	LogFileName    string
	LogEnvironment string
	writer         io.Writer
}

func (c *Config) isValid() (bool, error) {
	if c.LogTo != CONSOLE && c.LogTo != FILE {
		return false, fmt.Errorf("invalid LogTo value")
	}
	if c.OutputFormat != OUTPUT_FORMAT_JSON && c.OutputFormat != OUTPUT_FORMAT_TEXT {
		return false, fmt.Errorf("invalid OutputFormat value")
	}
	if c.LogEnvironment != DEVELOPMENT && c.LogEnvironment != STAGING && c.LogEnvironment != PRODUCTION {
		return false, fmt.Errorf("invalid LogEnvironment value")
	}
	return true, nil
}

func (c *Config) setupDefault() {
	c.LogEnvironment = DEVELOPMENT
	if os.Getenv("GOLOG_ENV") != "" {
		if os.Getenv("GOLOG_ENV") == DEVELOPMENT || os.Getenv("GOLOG_ENV") == STAGING || os.Getenv("GOLOG_ENV") == PRODUCTION {
			c.LogEnvironment = os.Getenv("GOLOG_ENV")
		}
	}
	c.LogTo = CONSOLE
	c.OutputFormat = OUTPUT_FORMAT_TEXT
	c.LogFileName = "golog.log"
	c.writer = os.Stdout
}

func (c *Config) configureLogTo(config *Config) error {
	if config.LogTo == FILE {
		f, err := os.OpenFile(c.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		c.writer = f
	}
	if config.LogTo == CONSOLE {
		c.writer = os.Stdout
	}
	return nil
}
