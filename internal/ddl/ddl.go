package ddl

type Database struct {
	Parameters *Parameters

	Schemas   []*Schema   `hcl:"schema,expand"`
	Functions []*Function `hcl:"function,expand"`
	Tables    []*Table    `hcl:"table,expand"`
	Triggers  []*Trigger  `hcl:"trigger,expand"`
}

type Parameters struct {
	SearchPath []string `hcl:"search_path"`
}

type Schema struct {
	Name string `hcl:",key"`
}

type Function struct {
	Schema string
	Name   string `hcl:",key"`

	Returns    string
	Language   string
	Definition string
}

type Table struct {
	Schema string
	Name   string `hcl:",key"`

	Columns []string
}

type Trigger struct {
	Schema string
	Name   string `hcl:",key"`

	When     string
	Events   []string
	Table    string
	ForEach  string `hcl:"for_each"`
	Function string
}
