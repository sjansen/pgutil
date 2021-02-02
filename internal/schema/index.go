package schema

import "strings"

// An Index is used to improve database performance
type Index struct {
	Schema string `hcl:"schema,label"`
	Table  string `hcl:"table,label"`
	Name   string `hcl:"name,label"`

	Primary bool   `hcl:"primary,optional"`
	Unique  bool   `hcl:"unique,optional"`
	Using   string `hcl:"using,optional"`
	Where   string `hcl:"where,optional"`

	Keys []*IndexKey `hcl:"key,block"`
}

// An IndexKey is used to improve database performance
type IndexKey struct {
	Column     string `hcl:"column,optional"`
	Expression string `hcl:"expression,optional"`
	OpClass    string `hcl:"opclass,optional"`
	Descending bool   `hcl:"descending,optional"`
}

func (idx *Index) ToSQL() (string, error) {
	var sb strings.Builder
	if idx.Primary {
		sb.WriteString("ALTER TABLE ")
		sb.WriteString(quoteName(idx.Table))
		sb.WriteString("\n  ADD ")
		if idx.Name != "" {
			sb.WriteString("CONSTRAINT ")
			sb.WriteString(quoteName(idx.Name))
			sb.WriteString("\n  ")
		}
		sb.WriteString("PRIMARY KEY (")
		for i, key := range idx.Keys {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(quoteName(key.Column))
		}
		sb.WriteString(")\n")
	}
	sb.WriteString(";\n")
	return sb.String(), nil
}
