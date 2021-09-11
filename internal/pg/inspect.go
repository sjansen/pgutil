package pg

import (
	"context"
	"sort"

	"github.com/sjansen/pgutil/internal/ddl"
)

type InspectOptions struct {
	SortChecks  bool
	SortColumns bool
	SortIndexes bool
}

// InspectDatabase describes the database
func (c *Conn) InspectDatabase(ctx context.Context, o *InspectOptions) (db *ddl.Database, err error) {
	db = &ddl.Database{}
	if o == nil {
		o = &InspectOptions{}
	}

	db.Parameters, err = c.ListParameters(ctx)
	if err != nil {
		return nil, err
	}

	db.Schemas, err = c.ListSchemas(ctx)
	if err != nil {
		return nil, err
	}

	db.Functions, err = c.ListFunctions(ctx)
	if err != nil {
		return nil, err
	}

	db.Sequences, err = c.ListSequences(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.inspectTables(ctx, o, db); err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Conn) inspectTables(ctx context.Context, o *InspectOptions, db *ddl.Database) (err error) {
	db.Tables, err = c.ListTables(ctx)
	if err != nil {
		return err
	}

	for _, table := range db.Tables {
		if err := c.inspectTable(ctx, o, db, table); err != nil {
			return err
		}
	}

	if o.SortIndexes {
		indexes := db.Indexes
		sort.Slice(indexes, func(i, j int) bool {
			return (indexes[i].Schema < indexes[j].Schema &&
				indexes[i].Table < indexes[j].Table &&
				indexes[i].Name < indexes[j].Name)

		})
	}

	return
}

func (c *Conn) inspectTable(ctx context.Context, o *InspectOptions, db *ddl.Database, table *ddl.Table) (err error) {
	table.Columns, err = c.ListColumns(ctx, table.Schema, table.Name)
	if err != nil {
		return err
	}
	if o.SortColumns {
		columns := table.Columns
		sort.Slice(columns, func(i, j int) bool {
			return columns[i].Name < columns[j].Name
		})
	}

	checks, err := c.ListChecks(ctx, table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(checks) > 0 {
		table.Checks = append(table.Checks, checks...)
	}
	if o.SortChecks {
		sort.Slice(checks, func(i, j int) bool {
			return (checks[i].Expression < checks[j].Expression)
		})
	}

	fks, err := c.ListForeignKeys(ctx, table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(fks) > 0 {
		table.ForeignKeys = append(table.ForeignKeys, fks...)
	}

	params, err := c.ListTableStorageParameters(ctx, table.Schema, table.Name)
	if err != nil {
		return err
	}
	table.StorageParameters = params

	indexes, err := c.ListIndexes(ctx, table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(indexes) > 0 {
		db.Indexes = append(db.Indexes, indexes...)
	}

	triggers, err := c.ListTriggers(ctx, table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(triggers) > 0 {
		db.Triggers = append(db.Triggers, triggers...)
	}

	return nil
}
