package ddl

import "strings"

func (t *Trigger) addColumn(s string) {
	if t.Columns == nil {
		t.Columns = []string{s}
	} else {
		t.Columns = append(t.Columns, s)
	}
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
