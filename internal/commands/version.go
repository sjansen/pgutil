// +build go1.12

package commands

import (
	"fmt"
	"io"
	"runtime/debug"
)

type VersionCmd struct {
	App   string
	Build string

	BuildInfo *debug.BuildInfo
	Verbose   bool
}

func (c *VersionCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	fmt.Fprintln(stdout, c.App, c.Build)
	if c.Verbose {
		fmt.Println("")
		for _, m := range c.BuildInfo.Deps {
			fmt.Printf("%-40s %-35s\n", m.Path, m.Version)
		}
	}
	return nil
}
