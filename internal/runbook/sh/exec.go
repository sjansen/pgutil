package sh

import (
	"context"
	"errors"
	"os/exec"
)

var _ execer = &Exec{}

type Exec struct {
	Args []string
}

func (x *Exec) exec(ctx context.Context, t *Target) error {
	cmd := exec.CommandContext(ctx, x.Args[0], x.Args[1:]...)
	cmd.Stdin = nil
	cmd.Stdout = t.stdout
	cmd.Stderr = t.stderr
	return cmd.Run()
}

func (x *Exec) Check() error {
	if len(x.Args) < 1 {
		return errors.New("too few args")
	}
	return nil
}
