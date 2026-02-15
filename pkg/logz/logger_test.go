package logz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/logz"
)

type TestStruct struct{}

func (ts TestStruct) getLogger() logz.Interface {
	return logz.Logger()
}

var testStruct TestStruct

type TestLogger struct{}

func (tl TestLogger) Infof(msg string, args ...interface{})  {}
func (tl TestLogger) Debugf(msg string, args ...interface{}) {}

var testLogger TestLogger

func TestLoggerForCaller(t *testing.T) {
	l := logz.Logger()
	dl := l.(logz.DefaultLogger)
	assert.Equal(t, "github.com/tartale/go/pkg/logz_test", dl.Name)

	l = testStruct.getLogger()
	dl = l.(logz.DefaultLogger)
	assert.Equal(t, "github.com/tartale/go/pkg/logz_test", dl.Name)
}

func TestLoggerForName(t *testing.T) {
	l := logz.LoggerForName("foo")
	dl := l.(logz.DefaultLogger)
	assert.Equal(t, "foo", dl.Name)
}

func TestLoggerForObjectTypePackagePath(t *testing.T) {
	l := logz.LoggerForObjectTypePackagePath(testStruct)
	dl := l.(logz.DefaultLogger)
	assert.Equal(t, "github.com/tartale/go/pkg/logz_test", dl.Name)
}

func TestSetLoggerForName(t *testing.T) {
	logz.SetLoggerForName("foo", testLogger)
	l := logz.LoggerForName("foo")
	dl := l.(TestLogger)
	assert.Equal(t, testLogger, dl)
}

func TestSetLoggerForCallerPackagePath(t *testing.T) {
	logz.SetLoggerForCallerPackagePath(testLogger)
	l := logz.Logger()
	dl := l.(TestLogger)
	assert.Equal(t, testLogger, dl)
}

func TestSetLoggerForObjectTypePackagePath(t *testing.T) {
	logz.SetLoggerForObjectTypePackagePath(testStruct, testLogger)
	l := logz.LoggerForObjectTypePackagePath(testStruct)
	dl := l.(TestLogger)
	assert.Equal(t, testLogger, dl)
}
