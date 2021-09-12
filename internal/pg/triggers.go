package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// DescribeTrigger describes a specific database table trigger
func (c *Conn) DescribeTrigger(ctx context.Context, schema, table, name string) (*ddl.Trigger, error) {
	db := pg10.New(c.conn)
	return db.DescribeTrigger(ctx, schema, table, name)
}

// ListTriggers describes the triggers of a database table
func (c *Conn) ListTriggers(ctx context.Context, schema, table string) ([]*ddl.Trigger, error) {
	db := pg10.New(c.conn)
	return db.ListTriggers(ctx, schema, table)
}
