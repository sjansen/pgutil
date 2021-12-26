package ddl

import (
	"fmt"
	"io"
	"time"

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
	Timestamp     time.Time
}

// WriteHCL converts structs describing a database to an an HCL configuration file.
func (db *Database) WriteHCL(w io.Writer, m *DatabaseMetadata) error {
	fmt.Fprint(w,
		"# Created: ", m.Timestamp.Local().Format(time.RFC822), "\n",
		"# Database: ", m.Database, "\n",
		"# Hostname: ", m.Host, "\n",
		"# Version: ", m.ServerVersion, "\n",
	)

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(db, f.Body())
	_, err := w.Write(f.Bytes())

	return err
}
