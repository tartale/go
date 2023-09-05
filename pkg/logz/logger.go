package logz

import "fmt"

type logger interface {
	Debugf(msg string, args ...interface{})
}

type defaultLogger struct{}

func (l defaultLogger) Debugf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

var Logger logger = defaultLogger{}
