package process

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Process struct {
	args []string
}

func Create(args []string) *Process {
	return &Process{args: args}
}

func (p *Process) Run(stdout, stderr io.Writer) error {
	arg0 := p.args[0]
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
		Args: p.args,

		Stdout: stdout,
		Stderr: stderr,
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
