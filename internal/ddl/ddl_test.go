package ddl_test

import (
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/ddl"
)

var configFile = `
parameters {
  search_path = ["$user", "public"]
}

schema "public" {}

function "update_modified_column" {
  returns = "trigger"
  language = "plpgsql"
  definition = <<EOF
BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
EOF
}

table "foo" {
  columns = ["id", "created", "modified", "key", "value"]
}

trigger "update_foo_modified" {
  when = "before"
  events = ["update"]
  table = "foo"
  for_each = "row"
  function = "update_modified_column"
}
`

var updateModifiedColumn = `BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
`

func TestHCL(t *testing.T) {
	require := require.New(t)

	expected := &ddl.Database{
		Parameters: &ddl.Parameters{
			SearchPath: []string{"$user", "public"},
		},
		Schemas: []*ddl.Schema{
			{
				Name: "public",
			},
		},
		Functions: []*ddl.Function{
			{
				Name:       "update_modified_column",
				Returns:    "trigger",
				Language:   "plpgsql",
				Definition: updateModifiedColumn,
			},
		},
		Tables: []*ddl.Table{
			{
				Name: "foo",
				Columns: []string{
					"id", "created", "modified", "key", "value",
				},
			},
		},
		Triggers: []*ddl.Trigger{
			{
				Name:     "update_foo_modified",
				When:     "before",
				Events:   []string{"update"},
				Table:    "foo",
				ForEach:  "row",
				Function: "update_modified_column",
			},
		},
	}

	var actual *ddl.Database
	err := hcl.Decode(&actual, configFile)
	require.NoError(err)
	require.Equal(expected, actual)
}
