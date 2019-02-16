package commands

import "fmt"

type RunBookListCmd struct {
	File string
}

type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

func (c *RunBookRunCmd) Run() error {
	fmt.Println("run:", c.File)
	return nil
}

func (c *RunBookListCmd) Run() error {
	fmt.Println("list:", c.File)
	return nil
}
