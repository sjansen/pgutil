package ddl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseError(t *testing.T) {
	require := require.New(t)

	e := &parseError{
		data: "foo\nbar\nbaz",
	}
	for _, tc := range []struct {
		cs       int
		expected string
	}{
		{cs: 1, expected: "foo\n ^"},
		{cs: 5, expected: "bar\n ^"},
		{cs: 10, expected: "baz\n  ^"},
	} {
		e.cs = tc.cs
		actual := e.Error()
		require.Equal(tc.expected, actual)
	}
}
