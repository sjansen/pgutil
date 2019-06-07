package ddl

//go:generate ragel-go -G2 -o index_parser.go index_parser.rl
//go:generate ragel-go -G2 -o foreign_key_parser.go foreign_key_parser.rl
//go:generate ragel-go -G2 -o trigger_parser.go trigger_parser.rl

// Database describes a PostgreSQL database
type Database struct {
	Parameters *Parameters `hcl:"parameters,block"`

	Schemas   []*Schema   `hcl:"schema,block"`
	Functions []*Function `hcl:"function,block"`
	Tables    []*Table    `hcl:"table,block"`
	Indexes   []*Index    `hcl:"index,block"`
	Triggers  []*Trigger  `hcl:"trigger,block"`
}

// Parameters describes database-level configuration options
type Parameters struct {
	SearchPath []string `hcl:"search_path,optional"`
}

// A Schema is a database namespace
type Schema struct {
	Name    string `hcl:"name,label"`
	Comment string `hcl:"comment,optional"`
	Owner   string `hcl:"owner,optional"`
}

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

// A Table is a collection of similar data organized as rows
type Table struct {
	Schema  string `hcl:"schema,label"`
	Name    string `hcl:"name,label"`
	Comment string `hcl:"comment,optional"`
	Owner   string `hcl:"owner,optional"`

	Columns     []*Column     `hcl:"column,block"`
	ForeignKeys []*ForeignKey `hcl:"foreign_key,block"`
}

// A Column is a data field of a table
type Column struct {
	Name    string `hcl:"name,label"`
	Type    string `hcl:"type,attr"`
	NotNull bool   `hcl:"not_null,optional"`
	Default string `hcl:"default,optional"`
}

// A ForeignKey ensures referential integrity
type ForeignKey struct {
	Table      string   `hcl:"table,label"`
	Columns    []string `hcl:"columns,attr"`
	Referenced []string `hcl:"referenced,attr"`

	Match    string `hcl:"match,optional"`
	OnDelete string `hcl:"on_delete,optional"`
	OnUpdate string `hcl:"on_update,optional"`

	Deferrable        bool `hcl:"deferrable,optional"`
	InitiallyDeferred bool `hcl:"initially_deferred,optional"`
}

// An Index is used to enhance database performance
type Index struct {
	Schema string `hcl:"schema,label"`
	Table  string `hcl:"table,label"`
	Name   string `hcl:"name,label"`

	Primary bool   `hcl:"primary,optional"`
	Unique  bool   `hcl:"unique,optional"`
	Using   string `hcl:"using,optional"`
	Where   string `hcl:"where,optional"`

	Keys []*IndexKey `hcl:"key,block"`
}

// An IndexKey is used to enhance database performance
type IndexKey struct {
	Column     string `hcl:"column,optional"`
	Expression string `hcl:"expression,optional"`
	OpClass    string `hcl:"opclass,optional"`
	Descending bool   `hcl:"descending,optional"`
}

// A Trigger executes a function when a specific event happens
type Trigger struct {
	Schema string `hcl:"schema,label"`
	Table  string `hcl:"table,label"`
	Name   string `hcl:"name,label"`

	Function string `hcl:"function,attr"`
	When     string `hcl:"when,attr"`

	Constraint        bool `hcl:"constraint,optional"`
	Deferrable        bool `hcl:"deferrable,optional"`
	InitiallyDeferred bool `hcl:"initially_deferred,optional"`
	ForEachRow        bool `hcl:"for_each_row,optional"`

	Delete   bool     `hcl:"delete,optional"`
	Insert   bool     `hcl:"insert,optional"`
	Truncate bool     `hcl:"truncate,optional"`
	Update   bool     `hcl:"update,optional"`
	Columns  []string `hcl:"columns,optional"`
}
