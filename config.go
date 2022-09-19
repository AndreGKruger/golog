package golog

import (
	"fmt"
	"io"
	"os"
)

// Supported GOLOG_ENV values
const (
	CONFIG_ENV_DEVELOPMENT = "development"
	CONFIG_ENV_STAGING     = "staging"
	CONFIG_ENV_PRODUCTION  = "production"
)

const (
	// OutputFormat constants
	CONFIG_OUTPUT_FORMAT_JSON = iota + 1
	CONFIG_OUTPUT_FORMAT_TEXT
	// LogTo constants
	CONFIG_LOG_TO_CONSOLE
	CONFIG_LOG_TO_FILE
	// Supported LogLevels
	CONFIG_LOG_LEVEL_DEBUG
	CONFIG_LOG_LEVEL_INFO
	CONFIG_LOG_LEVEL_WARN
	CONFIG_LOG_LEVEL_ERROR
)

type Config struct {
	LogTo          uint
	OutputFormat   uint
	LogFileName    string
	LogEnvironment string
	writer         io.Writer
}

func (c *Config) isValid() (bool, error) {
	if c.LogTo != CONFIG_LOG_TO_CONSOLE && c.LogTo != CONFIG_LOG_TO_FILE {
		return false, fmt.Errorf("invalid LogTo value")
	}
	if c.OutputFormat != CONFIG_OUTPUT_FORMAT_JSON && c.OutputFormat != CONFIG_OUTPUT_FORMAT_TEXT {
		return false, fmt.Errorf("invalid OutputFormat value")
	}
	if c.LogEnvironment != CONFIG_ENV_DEVELOPMENT && c.LogEnvironment != CONFIG_ENV_STAGING && c.LogEnvironment != CONFIG_ENV_PRODUCTION {
		return false, fmt.Errorf("invalid LogEnvironment value")
	}
	return true, nil
}

func (c *Config) setupDefault() {
	c.LogEnvironment = CONFIG_ENV_DEVELOPMENT
	if os.Getenv("GOLOG_ENV") != "" {
		if os.Getenv("GOLOG_ENV") == CONFIG_ENV_DEVELOPMENT || os.Getenv("GOLOG_ENV") == CONFIG_ENV_STAGING || os.Getenv("GOLOG_ENV") == CONFIG_ENV_PRODUCTION {
			c.LogEnvironment = os.Getenv("GOLOG_ENV")
		}
	}
	c.LogTo = CONFIG_LOG_TO_CONSOLE
	c.OutputFormat = CONFIG_OUTPUT_FORMAT_TEXT
	c.LogFileName = "golog.log"
	c.writer = os.Stdout
}

func (c *Config) configureLogTo(config *Config) error {
	if config.LogTo == CONFIG_LOG_TO_FILE {
		f, err := os.OpenFile(c.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		c.writer = f
	}
	if config.LogTo == CONFIG_LOG_TO_CONSOLE {
		c.writer = os.Stdout
	}
	return nil
}
