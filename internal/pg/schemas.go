package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListSchemas describes the schemas in the database
func (c *Conn) ListSchemas(ctx context.Context) ([]*ddl.Schema, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListSchemas(ctx)
}
