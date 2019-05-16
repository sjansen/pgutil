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
  schema = "public"
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
  schema = "public"
  columns = ["id", "created", "modified", "key", "value"]
}

trigger "update_foo_modified" {
  schema = "public"
  table = "foo"
  called = "before"
  event "update" {}
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
				Schema:     "public",
				Name:       "update_modified_column",
				Returns:    "trigger",
				Language:   "plpgsql",
				Definition: updateModifiedColumn,
			},
		},
		Tables: []*ddl.Table{
			{
				Schema: "public",
				Name:   "foo",
				Columns: []string{
					"id", "created", "modified", "key", "value",
				},
			},
		},
		Triggers: []*ddl.Trigger{
			{
				Schema: "public",
				Table:  "foo",
				Name:   "update_foo_modified",
				Called: "before",
				Events: []*ddl.TriggerEvent{
					{Event: "update"},
				},
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
