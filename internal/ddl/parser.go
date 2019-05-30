package ddl

import (
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
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
