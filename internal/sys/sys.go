package sys

import (
	"io"

	"go.uber.org/zap"
)

// IO encapsulates side effects
type IO struct {
	Log *zap.SugaredLogger

	Stdout io.Writer
	Stderr io.Writer
}
