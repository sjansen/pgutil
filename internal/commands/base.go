package commands

import (
	"os"

	"github.com/sjansen/pgutil/internal/sys"
)

type Base struct {
	sys.IO

	Debug     *os.File
	Verbosity int
}
