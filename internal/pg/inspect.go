package pg

import (
	"github.com/sjansen/pgutil/internal/ddl"
)

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
	}

	return db, nil
}
