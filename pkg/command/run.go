package command

import (
	"errors"
	"os"
	"os/exec"

	"github.com/tartale/go/pkg/logz"
)

// ErrDryRun is returned when DRY_RUN is enabled and a command is not executed.
var ErrDryRun = errors.New("dry run")

// StartIf starts cmd unless the DRY_RUN environment variable is set to "true",
// in which case it logs the command and returns ErrDryRun.
func StartIf(cmd *exec.Cmd) error {

	if os.Getenv("DRY_RUN") == "true" {
		logz.Logger().Infof("DRY RUN:\n\t%\n", cmd.String())
		return ErrDryRun
	}

	return cmd.Start()
}

// RunIf runs cmd unless the DRY_RUN environment variable is set to "true",
// in which case it logs the command and returns ErrDryRun.
func RunIf(cmd *exec.Cmd) error {

	if os.Getenv("DRY_RUN") == "true" {
		logz.Logger().Infof("DRY RUN:\n\t%s\n", cmd.String())
		return ErrDryRun
	}

	return cmd.Run()
}
