package ddl

import "strings"

func (t *Trigger) addColumn(s string) {
	t.Columns = append(t.Columns, s)
}

func (t *Trigger) setFunction(s string) {
	t.Function = s
}

func (t *Trigger) setName(name string) {
	t.Name = name
}

func (t *Trigger) setTable(table string) {
	t.Table = table
}

func (t *Trigger) setWhen(s string) {
	t.When = strings.ToUpper(
		collapseWhitespace(s),
	)
}

func (t *Trigger) ToSQL() (string, error) {
	var sb strings.Builder
	sb.WriteString("CREATE TRIGGER ")
	sb.WriteString(quoteName(t.Name))
	sb.WriteString("\n  ")
	sb.WriteString(t.When)
	if t.Update {
		sb.WriteString(" UPDATE")
	}
	sb.WriteString(" ON ")
	sb.WriteString(quoteName(t.Table))
	if t.ForEachRow {
		sb.WriteString("\n  FOR EACH ROW")
	}
	sb.WriteString("\n  EXECUTE PROCEDURE ")
	sb.WriteString(t.Function)
	sb.WriteString("()\n;\n")
	return sb.String(), nil
}
