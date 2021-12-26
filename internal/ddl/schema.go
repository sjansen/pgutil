package ddl

// A Schema is a database namespace
type Schema struct {
	Name    string `hcl:"name,label"`
	Owner   string `hcl:"owner,optional"`
	Comment string `hcl:"comment,optional"`
}
