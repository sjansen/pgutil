package catalog

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/catalog/pg96"
	"github.com/sjansen/pgutil/internal/ddl"
)

type Catalog interface {
	DescribeFunction(ctx context.Context, schema, name string) (*ddl.Function, error)
	DescribeTrigger(ctx context.Context, schema, table, name string) (*ddl.Trigger, error)
	ListChecks(ctx context.Context, schema, table string) ([]*ddl.Check, error)
	ListColumns(ctx context.Context, schema, table string) ([]*ddl.Column, error)
	ListForeignKeys(ctx context.Context, schema, table string) ([]*ddl.ForeignKey, error)
	ListFunctions(ctx context.Context) ([]*ddl.Function, error)
	ListIndexes(ctx context.Context, schema, table string) ([]*ddl.Index, error)
	ListSchemas(ctx context.Context) ([]*ddl.Schema, error)
	ListSequences(ctx context.Context) ([]*ddl.Sequence, error)
	ListTables(ctx context.Context) ([]*ddl.Table, error)
	ListTableStorageParameters(ctx context.Context, schema, table string) (*ddl.TableStorageParameters, error)
	ListTriggers(ctx context.Context, schema, table string) ([]*ddl.Trigger, error)
}

func New(ctx context.Context, conn *pgx.Conn) (Catalog, error) {
	var version string
	if err := conn.QueryRow(ctx, "SHOW server_version;").Scan(&version); err != nil {
		return nil, err
	}

	switch {
	case strings.HasPrefix(version, "9.6"):
		return pg96.New(conn), nil
	default:
		return pg10.New(conn), nil
	}
}
