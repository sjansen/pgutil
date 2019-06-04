package ddl

func (t *Index) addColumn(s string) {
	key := &IndexKey{Column: s}
	t.Keys = append(t.Keys, key)
}

func (t *Index) addExpression(s string) {
	key := &IndexKey{Expression: s}
	t.Keys = append(t.Keys, key)
}

func (t *Index) setName(name string) {
	t.Name = name
}

func (t *Index) setOpClass(opclass string) {
	key := t.Keys[len(t.Keys)-1]
	key.OpClass = opclass
}

func (t *Index) setTable(table string) {
	t.Table = table
}

func (t *Index) setUsing(s string) {
	t.Using = collapseWhitespace(s)
}

func (t *Index) setWhere(s string) {
	t.Where = collapseWhitespace(trim(s))
}
