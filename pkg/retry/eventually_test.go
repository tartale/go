package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventually_SucceedsBeforeTimeout(t *testing.T) {
	var (
		calls int
		want  = errors.New("temporary error")
	)

	err := Eventually(func() error {
		calls++
		if calls < 3 {
			return want
		}
		return nil
	}, 200*time.Millisecond, 5*time.Millisecond)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, calls, 3)
}

func TestEventually_TimesOutWhenErrorPersists(t *testing.T) {
	want := errors.New("always failing")

	err := Eventually(func() error {
		return want
	}, 30*time.Millisecond, 5*time.Millisecond)

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "timed out after")
		assert.ErrorIs(t, err, want)
	}
}

