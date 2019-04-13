package queues

import (
	"context"
	"io"
)

type Task interface {
	Start(ctx context.Context, stdout, stderr io.Writer)
}
