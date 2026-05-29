package testingx

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadersEqual(t *testing.T, expected, actual io.Reader) {
	expectedBytes := make([]byte, 512)
	actualBytes := make([]byte, 512)
	count := 0
	for {
		numExpectedBytes, err := expected.Read(expectedBytes)
		assert.NoError(t, err)
		numActualBytes, err := actual.Read(actualBytes)
		assert.NoError(t, err)
		assert.Equal(t, numExpectedBytes, numActualBytes,
			fmt.Sprintf("Streams differ at block starting at byte %d", count))
		assert.Equal(t, expectedBytes[:numExpectedBytes], actualBytes[:numActualBytes],
			fmt.Sprintf("Streams differ at block starting at byte %d", count))
		if numExpectedBytes < 512 || numActualBytes < 512 {
			break
		}
		count += 512
	}
}
