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
	sb.WriteString("CREATE ")
	if t.Constraint {
		sb.WriteString("CONSTRAINT ")
	}
	sb.WriteString("TRIGGER ")
	sb.WriteString(quoteName(t.Name))
	sb.WriteString("\n  ")
	sb.WriteString(t.Timing)
	switch {
	case t.Delete:
		sb.WriteString(" DELETE")
	case t.Insert:
		sb.WriteString(" INSERT")
	case t.Truncate:
		sb.WriteString(" TRUNCATE")
	case t.Update:
		sb.WriteString(" UPDATE")
	}
	sb.WriteString(" ON ")
	sb.WriteString(quoteName(t.Table))
	if t.Deferrable {
		sb.WriteString("\n  DEFERRABLE")
		if t.InitiallyDeferred {
			sb.WriteString(" INITIALLY DEFERRED")
		}
	}
	if t.ForEachRow {
		sb.WriteString("\n  FOR EACH ROW")
	} else {
		sb.WriteString("\n  FOR EACH STATEMENT")
	}
	sb.WriteString("\n  EXECUTE FUNCTION ")
	sb.WriteString(t.Function)
	sb.WriteString("()\n;\n")
	return sb.String(), nil
}
