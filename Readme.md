# Golog

## Introduction
Golog is a very simple logging library for Go. It is designed to be simple to use and to be easy to integrate into existing projects.

## Installation
To install Golog, simply run the following command:

    go get github.com/AndreGKruger/golog

## Usage
To use Golog, simply import the package into your project:

    import "github.com/AndreGKruger/golog"

Then, to create a new logger, simply call the `New` function:

    logger := golog.New()

To log a debug message, simply call the `Debug` function:

    logger.Debug("This is a debug message")

To log an info message, simply call the `Info` function:
    
    logger.Info("Hello, world!")

To log a warning message, simply call the `Warn` function:

    logger.Warn("Hello, world!")

To log an error message, simply call the `Error` function:
    
    logger.Error("Hello, world!")

## Configuration
Golog can be configured to log to a file, to log to the console, or both. To configure Golog, simply call the `Configure` function:

    logger.Configure(golog.Config{
        LogToFile: true,
        LogToConsole: true,
        LogFile: "log.txt",
    })

Output Format can be configured to be json or plain text. To configure Golog, simply call the `Configure` function:

    logger.Configure(golog.Config{
        OutputFormat: golog.OutputFormatJSON,
    })

## Environment logging
By default Golog will use the environment variable `GOLOG_ENV` to determine the log level. The following values are supported:
    
    GOLOG_ENV=development
    GOLOG_ENV=staging
    GOLOG_ENV=production

To override the environment variable, simply call the `SetEnvironment` function:
    logger.SetEnvironment(golog.DEVELOPMENT)

## Log Levels
Golog supports the following log levels:

    DEBUG
    INFO
    WARN
    ERROR

By default the following matrix of log levels and environments is used:

    DEBUG: development
    INFO: development, staging
    WARN: development, staging, production
    ERROR: development, staging, production

<!-- To override the log levels for a specific environment, simply call the `SetLogLevel` function:

    logger.SetLogLevel(golog.DEVELOPMENT, golog.DEBUG) -->

## Contributing
If you would like to contribute to Golog, please fork the repository and submit a pull request.

## License 
GNU General Public License v3.0 or later
A copy of the license is available in the repository's `LICENSE` file.

