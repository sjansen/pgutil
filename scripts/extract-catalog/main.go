package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
)

const getCatalogTablesSQL = `
SELECT
  c.table_name::text
, ARRAY_AGG(
    ARRAY[
      c.column_name::text
    , CASE WHEN c.data_type = 'ARRAY'
        THEN e.data_type||'[]'
        ELSE c.data_type
      END
    , CASE WHEN c.is_nullable = 'YES'
        THEN ''
	ELSE 'NOT NULL'
      END
    ] ORDER BY c.ordinal_position
  )
FROM information_schema.columns AS c
LEFT JOIN information_schema.element_types AS e
  ON (
    (c.table_catalog, c.table_schema, c.table_name, 'TABLE', c.dtd_identifier)
  = (e.object_catalog, e.object_schema, e.object_name, e.object_type, e.collection_type_identifier)
  )
WHERE c.table_schema = 'pg_catalog'
GROUP BY c.table_name
ORDER BY c.table_name
;
`

// IMPORTANT: New types should be added to sqlc.yaml overrides
var expectedColumnTypes = map[string]struct{}{
	`"char"`:                   {},
	`aclitem`:                  {},
	`anyarray`:                 {},
	`boolean`:                  {},
	`bytea`:                    {},
	`double precision`:         {},
	`integer`:                  {},
	`name`:                     {},
	`oid`:                      {},
	`pg_node_tree`:             {},
	`real`:                     {},
	`regproc`:                  {},
	`regtype`:                  {},
	`smallint`:                 {},
	`text`:                     {},
	`timestamp with time zone`: {},
	`xid`:                      {},
}

type Result struct {
	TableName string
	Columns   [][3]string
}

type Table struct {
	Name    string
	Columns []Column

	ColumnNameWidth int
	ColumnTypeWidth int
}

type Column struct {
	Name       string
	Type       string
	IsNullable bool
}

func getCatalogTables(ctx context.Context, conn *pgx.Conn) ([]*Table, error) {
	rows, err := conn.Query(ctx, getCatalogTablesSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := []*Table{}
	for rows.Next() {
		result := Result{}
		err := rows.Scan(&result.TableName, &result.Columns)
		if err != nil {
			return nil, err
		}

		table := &Table{
			Name:    result.TableName,
			Columns: make([]Column, 0, len(result.Columns)),
		}
		for _, column := range result.Columns {
			if table.ColumnNameWidth < len(column[0]) {
				table.ColumnNameWidth = len(column[0])
			}
			columnType := column[1]
			if table.ColumnTypeWidth < len(columnType) {
				table.ColumnTypeWidth = len(columnType)
			}
			if _, ok := expectedColumnTypes[columnType]; !ok {
				if strings.HasSuffix(columnType, "[]") {
					if _, ok := expectedColumnTypes[columnType[:len(columnType)-2]]; !ok {
						fmt.Fprintln(os.Stderr, "Unexpected column type:", columnType)
						os.Exit(1) // nolint
					}
				}
			}
			table.Columns = append(table.Columns, Column{
				Name:       column[0],
				Type:       column[1],
				IsNullable: len(column[2]) < 1,
			})
		}
		tables = append(tables, table)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

func listCatalogTables(tables []*Table) {
	for _, table := range tables {
		fmt.Printf("CREATE TABLE %s (\n", table.Name)
		for i, column := range table.Columns {
			if i < 1 {
				fmt.Print("    ")
			} else {
				fmt.Print(",   ")
			}
			if column.IsNullable {
				fmt.Printf("%-*s %s\n",
					table.ColumnNameWidth, column.Name, column.Type,
				)
			} else {
				fmt.Printf("%-*s %-*s NOT NULL\n",
					table.ColumnNameWidth, column.Name,
					table.ColumnTypeWidth, column.Type,
				)
			}
		}
		fmt.Println(")\n;")
	}
}

func showVersion(ctx context.Context, conn *pgx.Conn) error {
	var version string
	if err := conn.QueryRow(ctx, "SELECT version()").Scan(&version); err != nil {
		return err
	}

	fmt.Println("--", version)
	return nil
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	tables, err := getCatalogTables(ctx, conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = showVersion(ctx, conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	listCatalogTables(tables)
}
