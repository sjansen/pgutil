package ddl

import "strings"

// A Table is a collection of similar data organized as rows
type Table struct {
	Schema  string `hcl:"schema,label"`
	Name    string `hcl:"name,label"`
	Owner   string `hcl:"owner,optional"`
	Comment string `hcl:"comment,optional"`

	Columns     []*Column     `hcl:"column,block"`
	Checks      []*Check      `hcl:"check,block"`
	ForeignKeys []*ForeignKey `hcl:"foreign_key,block"`

	StorageParameters *TableStorageParameters `hcl:"storage_parameters,block"`
}

// A Column is a data field of a table
type Column struct {
	Name    string `hcl:"name,label"`
	Type    string `hcl:"type,attr"`
	NotNull bool   `hcl:"not_null,optional"`
	Default string `hcl:"default,optional"`
}

// A Check constraint limits column values
type Check struct {
	Name              string `hcl:"name,optional"`
	Expression        string `hcl:"expression,attr"`
	Deferrable        bool   `hcl:"deferrable,optional"`
	InitiallyDeferred bool   `hcl:"initially_deferred,optional"`
}

func (t *Table) ToSQL() (string, error) {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	sb.WriteString(quoteName(t.Name))
	sb.WriteString(" (\n")
	for i, col := range t.Columns {
		if i < 1 {
			sb.WriteString("  ")
		} else {
			sb.WriteString(", ")
		}
		sb.WriteString(quoteName(col.Name))
		sb.WriteString(" ")
		sb.WriteString(col.Type)
		if col.NotNull {
			sb.WriteString(" NOT NULL")
		}
		if col.Default != "" {
			sb.WriteString(" DEFAULT ")
			sb.WriteString(col.Default)
		}
		sb.WriteString("\n")
	}
	for _, check := range t.Checks {
		sb.WriteString(", CHECK(")
		sb.WriteString(check.Expression)
		sb.WriteString(")\n")
	}
	sb.WriteString(")\n;\n")
	return sb.String(), nil
}
