package ddl2

import "github.com/hashicorp/hcl2/hclparse"
import "github.com/hashicorp/hcl2/gohcl"

type Database struct {
	Tables []*Table `hcl:"table,block"`
}

type Table struct {
	Schema string `hcl:"schema,label"`
	Name   string `hcl:"name,label"`

	Comment string `hcl:"comment,attr"`
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
