package sh

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var _ execer = &Exec{}

// Exec describes how to run a command
type Exec struct {
	// Args specifies the command
	Args []string
	// Env specifies the environment of the command
	Env Env
}

func (x *Exec) exec(ctx context.Context, t *Target) error {
	cmd := exec.CommandContext(ctx, x.Args[0], x.Args[1:]...)
	cmd.Dir = t.basedir
	cmd.Env = x.Env.Apply(t.environ)
	cmd.Stdin = nil
	cmd.Stdout = t.stdout
	cmd.Stderr = t.stderr
	return cmd.Run()
}

func (x *Exec) Check() error {
	if len(x.Args) < 1 {
		return errors.New("too few args")
	}
	return validateArg0(x.Args[0])
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
