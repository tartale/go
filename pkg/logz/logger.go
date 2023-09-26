package logz

import (
	"fmt"

	"github.com/puzpuzpuz/xsync"
	"github.com/tartale/go/pkg/reflectx"
)

var (
	loggerMap = xsync.NewMapOf[Interface]()
)

type Interface interface {
	Infof(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
}

type DefaultLogger struct {
	Name string
}

func (l DefaultLogger) Infof(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (l DefaultLogger) Debugf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func Logger() Interface {
	packagePath := reflectx.CallerPackagePath(2)
	return LoggerForName(packagePath)
}

func LoggerForName(name string) Interface {
	l, _ := loggerMap.LoadOrStore(name, DefaultLogger{Name: name})
	return l
}

func LoggerForObjectTypePackagePath(obj any) Interface {
	objPackagePath := reflectx.ObjectTypePackagePath(obj)
	l, _ := loggerMap.LoadOrStore(objPackagePath, DefaultLogger{Name: objPackagePath})
	return l
}

func SetLoggerForName(name string, l Interface) {
	loggerMap.Store(name, l)
}

func SetLoggerForCallerPackagePath(l Interface) {
	callerPackagePath := reflectx.CallerPackagePath(2)
	loggerMap.Store(callerPackagePath, l)
}

func SetLoggerForObjectTypePackagePath(obj any, l Interface) {
	objectTypePackagePath := reflectx.ObjectTypePackagePath(obj)
	loggerMap.Store(objectTypePackagePath, l)
}
