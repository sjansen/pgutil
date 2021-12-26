package ddl

import (
	"fmt"
	"io"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclwrite"
)

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

// DatabaseMetadata describes a PostgreSQL database
type DatabaseMetadata struct {
	Host          string
	Database      string
	ServerVersion string
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

// WriteHCL converts structs describing a database to an an HCL configuration file.
func (db *Database) WriteHCL(w io.Writer, m *DatabaseMetadata) error {
	fmt.Fprint(w,
		"# Database: ", m.Database, "\n",
		"# Hostname: ", m.Host, "\n",
		"# Version: ", m.ServerVersion, "\n",
	)

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(db, f.Body())
	_, err := w.Write(f.Bytes())

	return err
}
