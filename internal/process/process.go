package process

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Command struct {
	Args []string

	Stdout io.Writer
	Stderr io.Writer
}

func (c *Command) Run() error {
	arg0 := c.Args[0]
	err := validateArg0(arg0)
	if err != nil {
		return err
	}

	command, err := exec.LookPath(arg0)
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{
		Path: command,
		Args: c.Args,

		Stdout: c.Stdout,
		Stderr: c.Stderr,
	}
	return cmd.Run()
}

func validateArg0(arg0 string) error {
	switch {
	case !strings.Contains(arg0, "/"):
	case strings.HasPrefix(arg0, "/"):
	case strings.HasPrefix(arg0, "./"):
	case strings.HasPrefix(arg0, "../"):
		break
	default:
		return fmt.Errorf("illegal command: %q", arg0)
	}
	return nil
}
