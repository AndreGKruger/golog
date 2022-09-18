package golog

import "fmt"

type Log interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Configure(Config)
}

type log struct{}

func New() Log {
	return &log{}
}

func (l *log) Debug(message string) {
	fmt.Println(message)
}

func (l *log) Info(message string) {
	fmt.Println(message)
}

func (l *log) Warn(message string) {
	fmt.Println(message)
}

func (l *log) Error(message string) {
	fmt.Println(message)
}

func (l *log) Configure(config Config) {
	// TODO
}
