package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListTableStorageParameters describes table storage parameters
func (c *Conn) ListTableStorageParameters(
	ctx context.Context, schema, table string,
) (*ddl.TableStorageParameters, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListTableStorageParameters(ctx, schema, table)
}
