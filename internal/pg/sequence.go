package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListSequences describes the sequences in the database
func (c *Conn) ListSequences(ctx context.Context) ([]*ddl.Sequence, error) {
	db := pg10.New(c.conn)
	return db.ListSequences(ctx)
}
