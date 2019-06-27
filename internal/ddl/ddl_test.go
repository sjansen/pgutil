package ddl_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/ddl"
)

var updateModifiedColumn = `BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
`

var expected = &ddl.Database{
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
			Name:    "bar",
			Comment: "",
			Columns: []*ddl.Column{
				{Name: "id", Type: "integer", NotNull: true},
				{Name: "foo_id", Type: "integer", NotNull: true},
			},
			ForeignKeys: []*ddl.ForeignKey{{
				Table:      "foo",
				Columns:    []string{"foo_id"},
				Referenced: []string{"id"},
			}},
		}, {
			Schema:  "public",
			Name:    "foo",
			Comment: "A simple test case",
			Columns: []*ddl.Column{
				{Name: "id", Type: "integer", NotNull: true},
				{Name: "created", Type: "timestamp with time zone", NotNull: true, Default: "now()"},
				{Name: "modified", Type: "timestamp with time zone", NotNull: true, Default: "now()"},
				{Name: "key", Type: "character varying(50)", NotNull: true},
				{Name: "value", Type: "character varying(500)"},
			},
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

func TestParseBytes(t *testing.T) {
	require := require.New(t)

	configFile, err := ioutil.ReadFile("testdata/example.hcl")
	require.NoError(err)

	actual, err := ddl.ParseBytes([]byte(configFile), "example.hcl")
	require.NoError(err)
	require.Equal(expected, actual)
}

func TestParseFile(t *testing.T) {
	require := require.New(t)

	actual, err := ddl.ParseFile("testdata/example.hcl")
	require.NoError(err)
	require.Equal(expected, actual)
}

func TestToSQL(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := ddl.ParseFile("testdata/settings.hcl")
	require.NoError(err)

	var buf bytes.Buffer
	for _, t := range db.Tables {
		sql, err := t.ToSQL()
		require.NoError(err)
		buf.WriteString(sql)
	}
	for _, idx := range db.Indexes {
		sql, err := idx.ToSQL()
		require.NoError(err)
		buf.WriteString(sql)
	}

	expected, err := ioutil.ReadFile("testdata/settings.sql")
	require.NoError(err)

	actual := buf.Bytes()
	if !assert.Equal(string(expected), string(actual)) {
		ioutil.WriteFile("testdata/settings-actual.sql", actual, 0666)
	}
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
