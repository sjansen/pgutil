package dispatcher

import (
	"context"

	"github.com/sjansen/pgutil/internal/dtos"
)

type Task interface {
	ID() string
	Dependencies() []string
	Run(ctx context.Context) *dtos.TaskStatus
}
