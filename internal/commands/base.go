package commands

import (
	"io"
	"os"

	"go.uber.org/zap"
)

type Base struct {
	Debug     *os.File
	Verbosity int

	Log    *zap.SugaredLogger
	Stdout io.Writer
	Stderr io.Writer
}
