package ddl

//go:generate ragel -Z -G2 -o check_parser.go check_parser.rl
//go:generate ragel -Z -G2 -o index_parser.go index_parser.rl
//go:generate ragel -Z -G2 -o foreign_key_parser.go foreign_key_parser.rl
//go:generate ragel -Z -G2 -o trigger_parser.go trigger_parser.rl

import (
	"io"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/hashicorp/hcl2/hclwrite"
)

// ParseBytes converts HCL data to structs describing a database.
// The filename is used for error messages.
func ParseBytes(src []byte, filename string) (*Database, error) {
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

// ParseFile converts an HCL configuration file to structs describing a database.
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

// Write converts structs describing a database to an an HCL configuration file.
func (db *Database) Write(w io.Writer) error {
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(db, f.Body())
	_, err := w.Write(f.Bytes())
	return err
}
