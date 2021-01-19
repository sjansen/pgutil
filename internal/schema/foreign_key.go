package schema

// A ForeignKey ensures referential integrity.
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
