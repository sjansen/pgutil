package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/taskset/parser"
)

func TestParseFile(t *testing.T) {
	require := require.New(t)

	_, err := parser.ParseFile("testdata/simple.hcl")
	require.NoError(err)
}
