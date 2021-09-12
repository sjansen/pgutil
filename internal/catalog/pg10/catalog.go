package pg10

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"

	"github.com/sjansen/pgutil/internal/ddl"
	"github.com/sjansen/pgutil/internal/sqlparser"
)

// DescribeFunction describes a custom function in the database
func (q *Queries) DescribeFunction(ctx context.Context, schema, name string) (*ddl.Function, error) {
	row, err := q.describeFunction(ctx, describeFunctionParams{
		Schema:       schema,
		FunctionName: name,
	})
	switch {
	case err == pgx.ErrNoRows:
		return nil, ErrNotFound
	case err != nil:
		return nil, err
	}

	fn := &ddl.Function{
		Schema:  schema,
		Name:    name,
		Owner:   row.Owner,
		Comment: row.Comment.String,

		Returns:    row.Returns.String,
		Language:   row.Language.String,
		Definition: row.Definition.String,
	}
	return fn, nil
}

// DescribeTrigger describes a specific database table trigger
func (q *Queries) DescribeTrigger(ctx context.Context, schema, table, name string) (*ddl.Trigger, error) {
	row, err := q.describeTrigger(ctx, describeTriggerParams{
		Schema:    schema,
		TableName: table,
		Trigger:   name,
	})
	switch {
	case err == pgx.ErrNoRows:
		return nil, ErrNotFound
	case err != nil:
		return nil, err
	}

	trigger, err := sqlparser.ParseCreateTrigger(row.String)
	if err != nil {
		return nil, err
	}
	return trigger, nil
}

// ListChecks describes the check constraints of a database table
func (q *Queries) ListChecks(ctx context.Context, schema, table string) ([]*ddl.Check, error) {
	rows, err := q.listChecks(ctx, listChecksParams{
		Schema:    schema,
		TableName: table,
	})
	if err != nil {
		return nil, err
	}

	var checks []*ddl.Check
	for _, row := range rows {
		check, err := sqlparser.ParseCheck(row.Definition.String)
		if err != nil {
			return nil, err
		}
		check.Name = row.Name
		checks = append(checks, check)
	}

	return checks, nil
}

// ListColumns describes the columns of a database table
func (q *Queries) ListColumns(ctx context.Context, schema, table string) ([]*ddl.Column, error) {
	rows, err := q.listColumns(ctx, listColumnsParams{
		Schema:    schema,
		TableName: table,
	})
	if err != nil {
		return nil, err
	}

	var columns []*ddl.Column
	for _, row := range rows {
		col := &ddl.Column{
			Name:    row.Name,
			Type:    row.Type.String,
			NotNull: row.NotNull,
			Default: row.Default.String,
			Comment: row.Comment.String,
		}
		columns = append(columns, col)
	}

	return columns, nil
}

// ListForeignKeys describes a database table's dependencies on other tables
func (q *Queries) ListForeignKeys(ctx context.Context, schema, table string) ([]*ddl.ForeignKey, error) {
	rows, err := q.listForeignKeys(ctx, listForeignKeysParams{
		Schema:    schema,
		TableName: table,
	})
	if err != nil {
		return nil, err
	}

	var fks []*ddl.ForeignKey
	for _, row := range rows {
		fk, err := sqlparser.ParseForeignKey(row.Definition.String)
		if err != nil {
			break
		}
		fk.Name = row.Name
		fks = append(fks, fk)
	}
	return fks, nil
}

// ListFunctions describes the custom functions in the database
func (q *Queries) ListFunctions(ctx context.Context) ([]*ddl.Function, error) {
	rows, err := q.listFunctions(ctx)
	if err != nil {
		return nil, err
	}

	var functions []*ddl.Function
	for _, row := range rows {
		fn := &ddl.Function{
			Schema:  row.Schema.String,
			Name:    row.Name,
			Owner:   row.Owner,
			Comment: row.Comment.String,

			Returns:    row.Returns.String,
			Language:   row.Language.String,
			Definition: row.Definition.String,
		}
		functions = append(functions, fn)
	}
	return functions, nil
}

// ListIndexes describes the indexes of a database table
func (q *Queries) ListIndexes(ctx context.Context, schema, table string) ([]*ddl.Index, error) {
	rows, err := q.listIndexes(ctx, listIndexesParams{
		Schema:    schema,
		TableName: table,
	})
	if err != nil {
		return nil, err
	}

	var indexes []*ddl.Index
	for _, row := range rows {
		index, err := sqlparser.ParseCreateIndex(row.Definition.String)
		if err != nil {
			return nil, err
		}
		index.Schema = schema
		index.Primary = row.IsPrimary
		indexes = append(indexes, index)
	}
	return indexes, nil
}

// ListSchemas describes the schemas in the database
func (q *Queries) ListSchemas(ctx context.Context) ([]*ddl.Schema, error) {
	rows, err := q.listSchemas(ctx)
	if err != nil {
		return nil, err
	}

	var schemas []*ddl.Schema
	for _, row := range rows {
		schema := &ddl.Schema{
			Name:    row.Name,
			Owner:   row.Owner,
			Comment: row.Comment.String,
		}
		schemas = append(schemas, schema)
	}
	return schemas, nil
}

// ListSequences describes the sequences in the database
func (q *Queries) ListSequences(ctx context.Context) ([]*ddl.Sequence, error) {
	rows, err := q.listSequences(ctx)
	if err != nil {
		return nil, err
	}

	var sequences []*ddl.Sequence
	for _, row := range rows {
		sequence := &ddl.Sequence{
			Schema:  row.Schema,
			Name:    row.Name,
			Owner:   row.Owner,
			Comment: row.Comment.String,

			Type:      row.Type.String,
			Start:     row.Start,
			Minimum:   row.Minimum,
			Maximum:   row.Maximum,
			Increment: row.Increment,
			Cache:     row.Cache,
			Cycle:     row.Cycle,
			OwnedBy: &ddl.SequenceOwner{
				Schema: row.OwnedBySchema,
				Table:  row.OwnedByTable,
				Column: row.OwnedByColumn,
			},
		}
		sequences = append(sequences, sequence)
	}
	return sequences, nil
}

// ListTables describes the tables in the database
func (q *Queries) ListTables(ctx context.Context) ([]*ddl.Table, error) {
	rows, err := q.listTables(ctx)
	if err != nil {
		return nil, err
	}

	var tables []*ddl.Table
	for _, row := range rows {
		table := &ddl.Table{
			Schema:  row.Schema.String,
			Name:    row.Name,
			Owner:   row.Owner,
			Comment: row.Comment.String,
		}
		tables = append(tables, table)
	}
	return tables, nil
}

// ListTableStorageParameters describes table storage parameters
func (q *Queries) ListTableStorageParameters(
	ctx context.Context, schema, table string,
) (*ddl.TableStorageParameters, error) {
	rows, err := q.listTableStorageParameters(ctx, listTableStorageParametersParams{
		Schema:    schema,
		TableName: table,
	})
	if err != nil {
		return nil, err
	}

	params := &ddl.TableStorageParameters{}
	for _, row := range rows {
		parts := strings.SplitN(row.String, "=", 2)
		if len(parts) < 2 {
			err = params.Set(parts[0], "")
		} else {
			err = params.Set(parts[0], parts[1])
		}
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}

// ListTriggers describes the triggers of a database table
func (q *Queries) ListTriggers(ctx context.Context, schema, table string) ([]*ddl.Trigger, error) {
	rows, err := q.listTriggers(ctx, listTriggersParams{
		Schema:    schema,
		TableName: table,
	})
	if err != nil {
		return nil, err
	}

	var triggers []*ddl.Trigger
	for _, row := range rows {
		trigger, err := sqlparser.ParseCreateTrigger(row.Definition.String)
		if err != nil {
			return nil, err
		}
		trigger.Schema = schema
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}
