package golog

import (
	"testing"
)

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
