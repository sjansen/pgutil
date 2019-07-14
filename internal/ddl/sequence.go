package ddl

// A Sequence generates a series of integer values
type Sequence struct {
	Schema  string `hcl:"schema,label"`
	Name    string `hcl:"name,label"`
	Owner   string `hcl:"owner,optional"`
	Comment string `hcl:"comment,optional"`

	DataType  string         `hcl:"data_type,optional"`
	Start     int64          `hcl:"start,optional"`
	Minimum   int64          `hcl:"minimum,optional"`
	Maximum   int64          `hcl:"maximum,optional"`
	Increment int64          `hcl:"increment,optional"`
	Cache     int64          `hcl:"cache,optional"`
	Cycle     bool           `hcl:"cycle,optional"`
	OwnedBy   *SequenceOwner `hcl:"owned_by,block"`
}

type SequenceOwner struct {
	Schema string `hcl:"schema,attr"`
	Table  string `hcl:"table,attr"`
	Column string `hcl:"column,attr"`
}
