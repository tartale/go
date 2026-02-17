package command

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunIf_DryRunSkipsExecution(t *testing.T) {
	t.Setenv("DRY_RUN", "true")

	cmd := exec.Command("definitely-not-a-real-binary")

	err := RunIf(cmd)

	assert.ErrorIs(t, err, ErrDryRun)
}

func TestStartIf_DryRunSkipsExecution(t *testing.T) {
	t.Setenv("DRY_RUN", "true")

	cmd := exec.Command("definitely-not-a-real-binary")

	err := StartIf(cmd)

	assert.ErrorIs(t, err, ErrDryRun)
}

func TestRunIf_ExecutesWhenNotDryRun(t *testing.T) {
	t.Setenv("DRY_RUN", "false")

	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess")
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")

	err := RunIf(cmd)

	assert.NoError(t, err)
}

// TestHelperProcess is a helper used by TestRunIf_ExecutesWhenNotDryRun. It is
// invoked as a subprocess and should not be run directly by `go test`.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	os.Exit(0)
}

