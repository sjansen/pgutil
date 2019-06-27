package ddl

import "strings"

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
