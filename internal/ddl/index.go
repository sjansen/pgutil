package ddl

func (t *Index) addColumn(s string) {
	key := &IndexKey{Column: s}
	t.Keys = append(t.Keys, key)
}

func (t *Index) setName(name string) {
	t.Name = name
}

func (t *Index) setTable(table string) {
	t.Table = table
}

func (t *Index) setUsing(s string) {
	t.Using = collapseWhitespace(s)
}
