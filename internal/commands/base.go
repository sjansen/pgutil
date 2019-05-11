package commands

import (
	"os"

	"github.com/sjansen/pgutil/internal/sys"
)

// Base contains dependencies and options common to all commands
type Base struct {
	sys.IO

	Debug     *os.File
	Verbosity int
}
