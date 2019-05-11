// +build !go1.12

package commands

import (
	"fmt"
)

// A VersionCmd is a request to reports the app's version
type VersionCmd struct {
	App   string
	Build string
}

// Run executes the command
func (c *VersionCmd) Run(base *Base) error {
	fmt.Fprintln(base.Stdout, c.App, c.Build)
	return nil
}
