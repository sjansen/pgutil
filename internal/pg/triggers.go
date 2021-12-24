package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// DescribeTrigger describes a specific database table trigger
func (c *Conn) DescribeTrigger(ctx context.Context, schema, table, name string) (*ddl.Trigger, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.DescribeTrigger(ctx, schema, table, name)
}

// ListTriggers describes the triggers of a database table
func (c *Conn) ListTriggers(ctx context.Context, schema, table string) ([]*ddl.Trigger, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListTriggers(ctx, schema, table)
}
