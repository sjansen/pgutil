package commands

import "fmt"

type VersionCmd struct {
	App   string
	Build string
}

func (c *VersionCmd) Run() error {
	fmt.Println(c.App, c.Build)
	return nil
}
