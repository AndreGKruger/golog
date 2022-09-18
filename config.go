package golog

const (
	OutputFormatJSON = iota + 1
	OutputFormatText
)

const (
	DEVELOPMENT = "development"
	STAGING     = "staging"
	PRODUCTION  = "production"
)

type Config struct {
	LogToFile      bool
	LogToConsole   bool
	LogFile        string
	LogEnvironment bool
	OutputFormat   uint
}
