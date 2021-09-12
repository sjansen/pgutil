package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// DescribeFunction describes a custom function in the database
func (c *Conn) DescribeFunction(ctx context.Context, schema, name string) (*ddl.Function, error) {
	db := pg10.New(c.conn)
	return db.DescribeFunction(ctx, schema, name)
}

// ListFunctions describes the custom functions in the database
func (c *Conn) ListFunctions(ctx context.Context) ([]*ddl.Function, error) {
	db := pg10.New(c.conn)
	return db.ListFunctions(ctx)
}
