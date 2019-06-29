package ddl

import "strings"

func (f *Function) ToSQL() (string, error) {
	var sb strings.Builder
	sb.WriteString("CREATE OR REPLACE FUNCTION ")
	sb.WriteString(f.Name)
	sb.WriteString("()\n")
	sb.WriteString("RETURNS ")
	sb.WriteString(f.Returns)
	sb.WriteString(" AS $$\n")
	sb.WriteString(f.Definition)
	sb.WriteString("$$ LANGUAGE '")
	sb.WriteString(f.Language)
	sb.WriteString("'\n;\n")
	return sb.String(), nil
}
