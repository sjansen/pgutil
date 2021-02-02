package ddl

// Database describes a PostgreSQL database
type Database struct {
	Parameters *Parameters `hcl:"parameters,block"`

	Schemas   []*Schema   `hcl:"schema,block"`
	Functions []*Function `hcl:"function,block"`
	Sequences []*Sequence `hcl:"sequence,block"`
	Tables    []*Table    `hcl:"table,block"`
	Indexes   []*Index    `hcl:"index,block"`
	Triggers  []*Trigger  `hcl:"trigger,block"`
}

// Parameters describes database-level configuration options
type Parameters struct {
	SearchPath []string `hcl:"search_path,optional"`
	Timezone   string   `hcl:"timezone,optional"`
}

// A Schema is a database namespace
type Schema struct {
	Name    string `hcl:"name,label"`
	Owner   string `hcl:"owner,optional"`
	Comment string `hcl:"comment,optional"`
}
