package ddl

// A ForeignKey ensures referential integrity
type ForeignKey struct {
	Name       string   `hcl:"name,optional"`
	Table      string   `hcl:"table,label"`
	Columns    []string `hcl:"columns,attr"`
	Referenced []string `hcl:"referenced,attr"`

	Match    string `hcl:"match,optional"`
	OnDelete string `hcl:"on_delete,optional"`
	OnUpdate string `hcl:"on_update,optional"`

	Deferrable        bool `hcl:"deferrable,optional"`
	InitiallyDeferred bool `hcl:"initially_deferred,optional"`
}

func (fk *ForeignKey) addColumn(s string) {
	fk.Columns = append(fk.Columns, s)
}

func (fk *ForeignKey) addReferenced(s string) {
	fk.Referenced = append(fk.Referenced, s)
}

func (fk *ForeignKey) setMatch(s string) {
	fk.Match = upper(s)
}

func (fk *ForeignKey) setOnDelete(s string) {
	fk.OnDelete = collapseWhitespace(upper(trim(s)))
}

func (fk *ForeignKey) setOnUpdate(s string) {
	fk.OnUpdate = collapseWhitespace(upper(trim(s)))
}

func (fk *ForeignKey) setTable(table string) {
	fk.Table = table
}
