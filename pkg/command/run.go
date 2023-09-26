package command

import (
	"errors"
	"os"
	"os/exec"

	"github.com/tartale/go/pkg/logz"
)

var ErrDryRun = errors.New("dry run")

func StartIf(cmd *exec.Cmd) error {

	if os.Getenv("DRY_RUN") == "true" {
		logz.Logger().Infof("DRY RUN:\n\t%\n", cmd.String())
		return ErrDryRun
	}

	return cmd.Start()
}

func RunIf(cmd *exec.Cmd) error {

	if os.Getenv("DRY_RUN") == "true" {
		logz.Logger().Infof("DRY RUN:\n\t%s\n", cmd.String())
		return ErrDryRun
	}

	return cmd.Run()
}
