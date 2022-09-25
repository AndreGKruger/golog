package golog

import (
	"testing"
)

// Test New
func TestNew(t *testing.T) {
	log := New()
	if log == nil {
		t.Error("Failed to create a new logger")
	}
}

// Test Default Configuration
func TestDefaultConfig(t *testing.T) {
	log := New()
	if log.(*golog).config.LogEnvironment != CONFIG_ENV_DEVELOPMENT {
		t.Error("Failed to set default configuration")
	}
	if log.(*golog).config.LogFileName != "golog.log" {
		t.Error("Failed to set default configuration")
	}
	if log.(*golog).config.LogTo != CONFIG_LOG_TO_CONSOLE {
		t.Error("Failed to set default configuration")
	}
	if log.(*golog).config.OutputFormat != CONFIG_OUTPUT_FORMAT_TEXT {
		t.Error("Failed to set default configuration")
	}
	if log.(*golog).config.writer == nil {
		t.Error("Failed to set default configuration")
	}
}

// Test shouldLog
func TestShouldLog(t *testing.T) {
	// Setup
	log := golog{
		config: &Config{},
		log:    nil,
	}
	log.config.setupDefault()

	// Test
	//Development
	got := log.shouldLog(CONFIG_LOG_LEVEL_DEBUG)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_INFO)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_WARN)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_ERROR)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}

	//Staging
	log.config.LogEnvironment = CONFIG_ENV_STAGING
	got = log.shouldLog(CONFIG_LOG_LEVEL_DEBUG)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_INFO)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_WARN)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_ERROR)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}

	//Production
	log.config.LogEnvironment = "production"
	got = log.shouldLog(CONFIG_LOG_LEVEL_DEBUG)
	if got != false {
		t.Errorf("shouldLog() = %v, want %v", got, false)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_INFO)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_WARN)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
	got = log.shouldLog(CONFIG_LOG_LEVEL_ERROR)
	if got != true {
		t.Errorf("shouldLog() = %v, want %v", got, true)
	}
}
