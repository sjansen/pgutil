package ddl2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/ddl2"
)

var configFile = `
table "public" "foo" {
  comment = "A simple test case"
}
`

func TestParse(t *testing.T) {
	require := require.New(t)

	expected := &ddl2.Database{
		Tables: []*ddl2.Table{
			{
				Schema:  "public",
				Name:    "foo",
				Comment: "A simple test case",
			},
		},
	}

	actual, err := ddl2.Parse([]byte(configFile), "foo.hcl")
	require.NoError(err)
	require.Equal(expected, actual)
}

func TestParseFile(t *testing.T) {
	require := require.New(t)

	expected := &ddl2.Database{
		Parameters: &ddl2.Parameters{
			SearchPath: []string{"$user", "public"},
		},
		Schemas: []*ddl2.Schema{
			{Name: "public"},
		},
		Functions: []*ddl2.Function{
			{
				Schema:   "public",
				Name:     "update_modified_column",
				Returns:  "trigger",
				Language: "plpgsql",
				Definition: `BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
`,
			},
		},

		Tables: []*ddl2.Table{
			{
				Schema:  "public",
				Name:    "foo",
				Comment: "A simple test case",
				Columns: []string{"id", "created", "modified", "key", "value"},
			},
		},
		Triggers: []*ddl2.Trigger{
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

	actual, err := ddl2.ParseFile("testdata/example.hcl")
	require.NoError(err)
	require.Equal(expected, actual)
}
