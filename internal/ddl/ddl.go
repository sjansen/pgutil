package ddl

//go:generate ragel-go -G2 -o trigger_parser.go trigger_parser.rl

// Database describes a PostgreSQL database
type Database struct {
	Parameters *Parameters

	Schemas   []*Schema   `hcl:"schema,expand"`
	Functions []*Function `hcl:"function,expand"`
	Tables    []*Table    `hcl:"table,expand"`
	Triggers  []*Trigger  `hcl:"trigger,expand"`
}

// Parameters describes database-level configuration options
type Parameters struct {
	SearchPath []string `hcl:"search_path"`
}

// A Schema is a database namespace
type Schema struct {
	Name    string `hcl:",key"`
	Owner   string
	Comment string
}

// A Function describes reusable behavior run on the server
type Function struct {
	Schema  string
	Name    string `hcl:",key"`
	Owner   string
	Comment string

	Returns    string
	Language   string
	Definition string
}

// A Table is a collection of similar data divided into rows
type Table struct {
	Schema  string
	Name    string `hcl:",key"`
	Owner   string
	Comment string

	Columns []string
}

// A Trigger executes a function when a specific event happens
type Trigger struct {
	Schema string
	Table  string
	Name   string `hcl:",key"`

	Function string
	When     string

	Constraint        bool
	Deferrable        bool
	InitiallyDeferred bool
	ForEachRow        bool `hcl:"for_each_row"`

	Delete   bool
	Insert   bool
	Truncate bool
	Update   bool
	Columns  []string
}
