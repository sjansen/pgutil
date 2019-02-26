// +build !go1.12

package commands

import (
	"fmt"
	"io"
)

type VersionCmd struct {
	App   string
	Build string
}

func (c *VersionCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	fmt.Fprintln(stdout, c.App, c.Build)
	return nil
}
