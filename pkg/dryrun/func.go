package dryrun

import (
	"os"
	"strings"

	"github.com/tartale/go/pkg/logz"
)

func Do[RET any](fn func() RET, msg ...string) RET {

	if os.Getenv("DRY_RUN") == "true" {
		if len(msg) > 0 {
			logz.Logger().Infof(strings.Join(msg, " ") + "\n")
		}
		var zero RET
		return zero
	}

	return fn()
}
