package logz

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/puzpuzpuz/xsync"
)

var (
	loggerMap = xsync.NewMapOf[logger]()
)

type logger interface {
	Debugf(msg string, args ...interface{})
}

type defaultLogger struct{}

func (l defaultLogger) Debugf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func SetLogger(packageName string, l logger) {
	loggerMap.Store(packageName, l)
}

func LoggerForPackage(packageName string) logger {
	l, _ := loggerMap.LoadOrStore(packageName, defaultLogger{})
	return l
}

func Logger() logger {
	if pc, _, _, ok := runtime.Caller(0); ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			packageName := strings.Split(fn.Name(), ".")[0]
			return LoggerForPackage(packageName)
		}
	}

	return defaultLogger{}
}
