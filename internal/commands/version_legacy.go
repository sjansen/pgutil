// +build !go1.12

package commands

import (
	"fmt"
)

type VersionCmd struct {
	App   string
	Build string
}

func (c *VersionCmd) Run(base *Base) error {
	fmt.Fprintln(base.Stdout, c.App, c.Build)
	return nil
}
