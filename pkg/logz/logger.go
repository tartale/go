package logz

import (
	"fmt"

	"github.com/puzpuzpuz/xsync"
	"github.com/tartale/go/pkg/reflectx"
)

var (
	loggerMap = xsync.NewMapOf[Interface]()
)

// Interface is the minimal logging interface used by this package.
type Interface interface {
	Infof(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
}

// DefaultLogger is a basic implementation of Interface that writes to stdout.
type DefaultLogger struct {
	Name string
}

// Infof logs a formatted info-level message.
func (l DefaultLogger) Infof(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

// Debugf logs a formatted debug-level message.
func (l DefaultLogger) Debugf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

// Logger returns a logger associated with the caller's package path.
//
// Example:
//
//	log := logz.Logger()
//	log.Infof("hello %s", "world")
func Logger() Interface {
	packagePath := reflectx.CallerPackagePath(2)
	return LoggerForName(packagePath)
}

// LoggerForName returns a logger registered under the given name,
// creating a DefaultLogger if one does not already exist.
func LoggerForName(name string) Interface {
	l, _ := loggerMap.LoadOrStore(name, DefaultLogger{Name: name})
	return l
}

// LoggerForObjectTypePackagePath returns a logger associated with the package
// path of obj's concrete type, creating a DefaultLogger if necessary.
func LoggerForObjectTypePackagePath(obj any) Interface {
	objPackagePath := reflectx.ObjectTypePackagePath(obj)
	l, _ := loggerMap.LoadOrStore(objPackagePath, DefaultLogger{Name: objPackagePath})
	return l
}

// SetLoggerForName registers l as the logger for the given name.
//
// Example:
//
//	logz.SetLoggerForName("my-app", myLogger)
func SetLoggerForName(name string, l Interface) {
	loggerMap.Store(name, l)
}

// SetLoggerForCallerPackagePath registers l as the logger for the caller's package path.
func SetLoggerForCallerPackagePath(l Interface) {
	callerPackagePath := reflectx.CallerPackagePath(2)
	loggerMap.Store(callerPackagePath, l)
}

// SetLoggerForObjectTypePackagePath registers l as the logger for the
// package path of obj's concrete type.
func SetLoggerForObjectTypePackagePath(obj any, l Interface) {
	objectTypePackagePath := reflectx.ObjectTypePackagePath(obj)
	loggerMap.Store(objectTypePackagePath, l)
}
