package ddl

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
