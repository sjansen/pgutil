package exec

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Task struct {
	Args   []string
	Stdout io.Writer
	Stderr io.Writer
}

func (t *Task) Start(ctx context.Context) error {
	arg0 := t.Args[0]
	err := validateArg0(arg0)
	if err != nil {
		return err
	}

	command, err := exec.LookPath(arg0)
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{
		Path:   command,
		Args:   t.Args,
		Stdout: t.Stdout,
		Stderr: t.Stderr,
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
