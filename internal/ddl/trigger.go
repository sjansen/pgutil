package ddl

import "strings"

// A Trigger executes a function when a specific event happens.
type Trigger struct {
	Schema string `hcl:"schema,label"`
	Table  string `hcl:"table,label"`
	Name   string `hcl:"name,label"`

	From     string `hcl:"from,optional"`
	Function string `hcl:"function,attr"`
	Timing   string `hcl:"timing,attr"`
	When     string `hcl:"when,optional"`

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

// ToSQL produces a CREATE TRIGGER statement.
func (t *Trigger) ToSQL() (string, error) {
	var sb strings.Builder
	sb.WriteString("CREATE TRIGGER ")
	sb.WriteString(quoteName(t.Name))
	sb.WriteString("\n  ")
	sb.WriteString(t.Timing)
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
