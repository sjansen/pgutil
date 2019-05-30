package ddl_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/ddl"
)

var configFile = `
table "public" "foo" {
  comment = "A simple test case"
  columns = ["id", "created", "modified", "key", "value"]
}
`

var updateModifiedColumn = `BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
`

func TestParseBytes(t *testing.T) {
	require := require.New(t)

	expected := &ddl.Database{
		Tables: []*ddl.Table{
			{
				Schema:  "public",
				Name:    "foo",
				Comment: "A simple test case",
				Columns: []string{"id", "created", "modified", "key", "value"},
			},
		},
	}

	actual, err := ddl.ParseBytes([]byte(configFile), "foo.hcl")
	require.NoError(err)
	require.Equal(expected, actual)
}

func TestParseFile(t *testing.T) {
	require := require.New(t)

	expected := &ddl.Database{
		Parameters: &ddl.Parameters{
			SearchPath: []string{"$user", "public"},
		},
		Schemas: []*ddl.Schema{
			{Name: "public"},
		},
		Functions: []*ddl.Function{
			{
				Schema:     "public",
				Name:       "update_modified_column",
				Returns:    "trigger",
				Language:   "plpgsql",
				Definition: updateModifiedColumn,
			},
		},

		Tables: []*ddl.Table{
			{
				Schema:  "public",
				Name:    "foo",
				Comment: "A simple test case",
				Columns: []string{"id", "created", "modified", "key", "value"},
			},
		},
		Triggers: []*ddl.Trigger{
			{
				Schema:     "public",
				Table:      "foo",
				Name:       "update_foo_modified",
				Function:   "update_modified_column",
				When:       "before",
				Update:     true,
				ForEachRow: true,
			},
		},
	}

	actual, err := ddl.ParseFile("testdata/example.hcl")
	require.NoError(err)
	require.Equal(expected, actual)
}
