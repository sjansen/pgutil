package ddl2

import "github.com/hashicorp/hcl2/hclparse"
import "github.com/hashicorp/hcl2/gohcl"

// Database describes a PostgreSQL database
type Database struct {
	Parameters *Parameters `hcl:"parameters,block"`

	Schemas   []*Schema   `hcl:"schema,block"`
	Functions []*Function `hcl:"function,block"`
	Tables    []*Table    `hcl:"table,block"`
	Triggers  []*Trigger  `hcl:"trigger,block"`
}

// Parameters describes database-level configuration options
type Parameters struct {
	SearchPath []string `hcl:"search_path,optional"`
}

// A Schema is a database namespace
type Schema struct {
	Name string `hcl:"name,label"`
}

// A Function describes reusable behavior run on the server
type Function struct {
	Schema  string `hcl:"schema,label"`
	Name    string `hcl:"name,label"`
	Comment string `hcl:"comment,optional"`

	Returns    string `hcl:"returns,attr"`
	Language   string `hcl:"language,attr"`
	Definition string `hcl:"definition,attr"`
}

// A Table is a collection of similar data organized as rows
type Table struct {
	Schema  string `hcl:"schema,label"`
	Name    string `hcl:"name,label"`
	Comment string `hcl:"comment,optional"`

	Columns []string `hcl:"columns,optional"`
}

// A Trigger executes a function when a specific event happens
type Trigger struct {
	Schema  string `hcl:"schema,label"`
	Table   string `hcl:"table,label"`
	Name    string `hcl:"name,label"`
	Comment string `hcl:"comment,optional"`

	Function   string `hcl:"function,attr"`
	When       string `hcl:"when,attr"`
	Update     bool   `hcl:"update,attr"`
	ForEachRow bool   `hcl:"for_each_row,attr"`
}

func Parse(src []byte, filename string) (*Database, error) {
	p := hclparse.NewParser()
	f, err := p.ParseHCL(src, filename)
	if err != nil {
		return nil, err
	}

	db := &Database{}
	err = gohcl.DecodeBody(f.Body, nil, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ParseFile(filename string) (*Database, error) {
	p := hclparse.NewParser()
	f, err := p.ParseHCLFile(filename)
	if err != nil {
		return nil, err
	}

	db := &Database{}
	err = gohcl.DecodeBody(f.Body, nil, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
