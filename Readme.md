# Golog

## Introduction
Golog is a very simple logging library for Go. It is designed to be simple to use and to be easy to integrate into existing projects.

---

## Installation
To install Golog, simply run the following command:

    go get github.com/AndreGKruger/golog@latest

---

## Usage
To use Golog, simply import the package into your project:

    import "github.com/AndreGKruger/golog"

To create a new logger, simply call the `New` function:

    logger := golog.New()

To log a debug message, simply call the `Debug`, `Info`, `Warn` or `Error` function:

    logger.Debug("This is a debug message")
    logger.Debug("This is a debug message with arguments", arg1, arg2, arg3)
    logger.Info("Hello, world!")
    logger.Info("Hello, world! with arguments", arg1, arg2, arg3)
    logger.Warn("Hello, world!")
    logger.Warn("Hello, world! with arguments", arg1, arg2, arg3)
    logger.Error("Hello, world!")
    logger.Error("Hello, world! with arguments", arg1, arg2, arg3)

---

## Example output
Output

    [ 2022/09/20 17:07:07 ] - [ENV:development] - [DEBUG]: sample log message  - [ARGS]: [{Test:Testing}]

---

## Configuration
Golog can be configured to log to a file, to log to the console. To configure Golog, simply call the `Configure` function:

    ok, err := logger.Configure(golog.Config{
        LogEnvironment: golog.CONFIG_ENV_DEVELOPMENT,
        LogFileName:    "log.txt",
        LogTo:          golog.CONFIG_LOG_TO_CONSOLE,
        OutputFormat:   golog.CONFIG_OUTPUT_FORMAT_TEXT,
    })

Json Output Format can be configured by passing the `OUTPUT_FORMAT_JSON` constant to the `OutputFormat` field of the `Config` struct:

    OutputFormat: golog.CONFIG_OUTPUT_FORMAT_JSON,

---

## Environment logging
By default Golog will use the environment variable `GOLOG_ENV` to determine the log level. The following values are supported:
    
    GOLOG_ENV=development
    GOLOG_ENV=staging
    GOLOG_ENV=production

The GOLOG_ENV will override the environment variable that is set during the `SetEnvironment` function:
    logger.SetEnvironment(golog.DEVELOPMENT)

---

## Log Levels
Golog supports the following log levels:

    CONFIG_LOG_LEVEL_DEBUG
    CONFIG_LOG_LEVEL_INFO
    CONFIG_LOG_LEVEL_WARN
    CONFIG_LOG_LEVEL_ERROR

By default the following matrix of log levels and environments is used:

    CONFIG_LOG_LEVEL_DEBUG: development
    CONFIG_LOG_LEVEL_INFO: development, staging
    CONFIG_LOG_LEVEL_WARN: development, staging, production
    CONFIG_LOG_LEVEL_ERROR: development, staging, production

---

## Contributing
If you would like to contribute to Golog, please fork the repository and submit a pull request.

---

## License 
GNU General Public License v3.0 or later
A copy of the license is available in the repository's `LICENSE` file.

---
