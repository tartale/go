package asserts

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadersEqual(t *testing.T, expected, actual io.Reader) bool {
	expectedBytes := make([]byte, 512)
	actualBytes := make([]byte, 512)
	count := 0
	for {
		numExpectedBytes, err := expected.Read(expectedBytes)
		assert.NoError(t, err)
		numActualBytes, err := actual.Read(actualBytes)
		assert.NoError(t, err)
		if !assert.ObjectsAreEqual(numExpectedBytes, numActualBytes) {
			return assert.Fail(t, fmt.Sprintf("Streams differ at block starting at byte %d", count))
		}
		if !assert.ObjectsAreEqual(expectedBytes[:numExpectedBytes], actualBytes[:numActualBytes]) {
			return assert.Fail(t, fmt.Sprintf("Streams differ at block starting at byte %d", count))
		}
		if numExpectedBytes < 512 || numActualBytes < 512 {
			break
		}
		count += 512
	}
	return true
}
