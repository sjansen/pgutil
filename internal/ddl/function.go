package ddl

import "strings"

// A Function describes reusable behavior run on the server
type Function struct {
	Schema  string `hcl:"schema,label"`
	Name    string `hcl:"name,label"`
	Comment string `hcl:"comment,optional"`
	Owner   string `hcl:"owner,optional"`

	Returns    string `hcl:"returns,attr"`
	Language   string `hcl:"language,attr"`
	Definition string `hcl:"definition,attr"`
}

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
