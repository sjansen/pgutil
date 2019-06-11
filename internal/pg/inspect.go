package pg

import (
	"sort"

	"github.com/sjansen/pgutil/internal/ddl"
)

type InspectOptions struct {
	SortChecks  bool
	SortColumns bool
	SortIndexes bool
}

// InspectDatabase describes the database
func (c *Conn) InspectDatabase(o *InspectOptions) (db *ddl.Database, err error) {
	db = &ddl.Database{}
	if o == nil {
		o = &InspectOptions{}
	}

	db.Schemas, err = c.ListSchemas()
	if err != nil {
		return nil, err
	}

	db.Functions, err = c.ListFunctions()
	if err != nil {
		return nil, err
	}

	db.Tables, err = c.ListTables()
	if err != nil {
		return nil, err
	}

	for _, table := range db.Tables {
		if err := c.inspectTable(o, db, table); err != nil {
			return nil, err
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

	return db, nil
}

func (c *Conn) inspectTable(o *InspectOptions, db *ddl.Database, table *ddl.Table) (err error) {
	table.Columns, err = c.ListColumns(table.Schema, table.Name)
	if err != nil {
		return err
	}
	if o.SortColumns {
		columns := table.Columns
		sort.Slice(columns, func(i, j int) bool {
			return columns[i].Name < columns[j].Name
		})
	}

	checks, err := c.ListChecks(table.Schema, table.Name)
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

	fks, err := c.ListForeignKeys(table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(fks) > 0 {
		table.ForeignKeys = append(table.ForeignKeys, fks...)
	}

	indexes, err := c.ListIndexes(table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(indexes) > 0 {
		db.Indexes = append(db.Indexes, indexes...)
	}

	triggers, err := c.ListTriggers(table.Schema, table.Name)
	if err != nil {
		return err
	}
	if len(triggers) > 0 {
		db.Triggers = append(db.Triggers, triggers...)
	}

	return nil
}
