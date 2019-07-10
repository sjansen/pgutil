package ddl

import "strings"

// A Trigger executes a function when a specific event happens
type Trigger struct {
	Schema string `hcl:"schema,label"`
	Table  string `hcl:"table,label"`
	Name   string `hcl:"name,label"`

	Function string `hcl:"function,attr"`
	When     string `hcl:"when,attr"`

	Constraint        bool `hcl:"constraint,optional"`
	Deferrable        bool `hcl:"deferrable,optional"`
	InitiallyDeferred bool `hcl:"initially_deferred,optional"`
	ForEachRow        bool `hcl:"for_each_row,optional"`

	Delete   bool     `hcl:"delete,optional"`
	Insert   bool     `hcl:"insert,optional"`
	Truncate bool     `hcl:"truncate,optional"`
	Update   bool     `hcl:"update,optional"`
	Columns  []string `hcl:"columns,optional"`
}

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
