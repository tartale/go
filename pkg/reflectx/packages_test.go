package reflectx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/reflectx"
)

type TestStruct struct{}

func (ts TestStruct) getPackagePath() string {
	return reflectx.CallerPackagePath(1)
}

var testStruct TestStruct

func TestCallerPackagePath(t *testing.T) {

	packagePath := reflectx.CallerPackagePath(1)
	assert.Equal(t, "github.com/tartale/go/pkg/reflectx_test", packagePath)

	packagePath = testStruct.getPackagePath()
	assert.Equal(t, "github.com/tartale/go/pkg/reflectx_test", packagePath)
}

func BenchmarkCallerPackagePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectx.CallerPackagePath(1)
	}
}

func TestObjectTypePackagePath(t *testing.T) {

	packagePath := reflectx.ObjectTypePackagePath(testStruct)
	assert.Equal(t, "github.com/tartale/go/pkg/reflectx_test", packagePath)
}

func BenchmarkObjectTypePackagePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectx.ObjectTypePackagePath(testStruct)
	}
}

func TestObjectTypePath(t *testing.T) {

	objPath := reflectx.ObjectTypePath(testStruct)
	assert.Equal(t, "github.com/tartale/go/pkg/reflectx_test.TestStruct", objPath)
}

func BenchmarkObjectTypePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectx.ObjectTypePath(testStruct)
	}
}
