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

func TestHCL(t *testing.T) {
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
