package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListTableStorageParameters describes table storage parameters
func (c *Conn) ListTableStorageParameters(
	ctx context.Context, schema, table string,
) (*ddl.TableStorageParameters, error) {
	db := pg10.New(c.conn)
	return db.ListTableStorageParameters(ctx, schema, table)
}
