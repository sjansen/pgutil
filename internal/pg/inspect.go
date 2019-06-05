package pg

import (
	"sort"

	"github.com/sjansen/pgutil/internal/ddl"
)

type InspectOptions struct {
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

	for i, table := range db.Tables {
		db.Tables[i].Columns, err = c.ListColumns(table.Schema, table.Name)
		if err != nil {
			return nil, err
		}
		if o.SortColumns {
			columns := db.Tables[i].Columns
			sort.Slice(columns, func(i, j int) bool {
				return columns[i].Name < columns[j].Name
			})
		}

		indexes, err := c.ListIndexes(table.Schema, table.Name)
		if err != nil {
			return nil, err
		}
		if len(indexes) > 0 {
			db.Indexes = append(db.Indexes, indexes...)
		}

		triggers, err := c.ListTriggers(table.Schema, table.Name)
		if err != nil {
			return nil, err
		}
		if len(triggers) > 0 {
			db.Triggers = append(db.Triggers, triggers...)
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
