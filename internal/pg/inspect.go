package pg

import (
	"github.com/sjansen/pgutil/internal/ddl"
)

// InspectDatabase describes the database
func (c *Conn) InspectDatabase() (db *ddl.Database, err error) {
	db = &ddl.Database{}

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

	return db, nil
}
