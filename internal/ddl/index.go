package ddl

import "strings"

func (idx *Index) addColumn(s string) {
	key := &IndexKey{Column: s}
	idx.Keys = append(idx.Keys, key)
}

func (idx *Index) addExpression(s string) {
	key := &IndexKey{Expression: s}
	idx.Keys = append(idx.Keys, key)
}

func (idx *Index) setName(name string) {
	idx.Name = name
}

func (idx *Index) setOpClass(opclass string) {
	key := idx.Keys[len(idx.Keys)-1]
	key.OpClass = opclass
}

func (idx *Index) setTable(table string) {
	idx.Table = table
}

func (idx *Index) setUsing(s string) {
	idx.Using = collapseWhitespace(s)
}

func (idx *Index) setWhere(s string) {
	idx.Where = collapseWhitespace(trim(s))
}

func (idx *Index) ToSQL() (string, error) {
	var sb strings.Builder
	if idx.Primary {
		sb.WriteString("ALTER TABLE ")
		sb.WriteString(quoteName(idx.Table))
		sb.WriteString("\n  ADD ")
		if idx.Name != "" {
			sb.WriteString("CONSTRAINT ")
			sb.WriteString(quoteName(idx.Name))
			sb.WriteString("\n  ")
		}
		sb.WriteString("PRIMARY KEY (")
		for i, key := range idx.Keys {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(quoteName(key.Column))
		}
		sb.WriteString(")\n")
	}
	sb.WriteString(";\n")
	return sb.String(), nil
}
