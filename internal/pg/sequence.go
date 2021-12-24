package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListSequences describes the sequences in the database
func (c *Conn) ListSequences(ctx context.Context) ([]*ddl.Sequence, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListSequences(ctx)
}
