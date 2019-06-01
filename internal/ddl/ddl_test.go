package ddl_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestWrite(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := ddl.ParseFile("testdata/example.hcl")
	require.NoError(err)

	expected, err := ioutil.ReadFile("testdata/expected.hcl")
	require.NoError(err)

	buf := &bytes.Buffer{}
	err = ddl.Write(buf, db)
	require.NoError(err)

	actual := buf.Bytes()
	if !assert.Equal(string(expected), string(actual)) {
		ioutil.WriteFile("testdata/actual.hcl", actual, 0666)
	}
}
