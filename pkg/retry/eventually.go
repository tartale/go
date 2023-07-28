package retry

import (
	"fmt"
	"time"
)

// Eventually takes a function, timeout and retry duration and
// executes the function repeatedly until it returns a nil error,
// or the timeout expires.
//
// Inspired by the gomega "Eventually" assertion
func Eventually(fn func() error, timeout, retryAfter time.Duration) error {
	timeoutAt := time.Now().Add(timeout)
	for {
		err := fn()
		if err == nil {
			return nil
		}
		now := time.Now()
		if now.After(timeoutAt) {
			return fmt.Errorf("timed out after %s; error: %w", timeout.String(), err)
		}
		time.Sleep(retryAfter)
	}
}
