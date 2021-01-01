package scanner_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/scanner"
)

func TestNewScanner(t *testing.T) {
	require := require.New(t)
	for _, tc := range []struct {
		Label string
		Input string
		Error string
	}{{
		Label: "invalid",
		Input: string([]byte{128}),
		Error: "invalid utf8",
	}} {
		_, err := scanner.New(tc.Input)
		if tc.Error == "" {
			require.Nil(err, tc.Label)
		} else {
			require.Error(err, tc.Label)
			require.Equal(tc.Error, err.Error(), tc.Label)
		}
	}
}

func TestPosition(t *testing.T) {
	require := require.New(t)
	s, err := scanner.New("EXPLAIN ANALYZE\r\nSELECT")
	require.NoError(err)

	err = s.SkipWhitespace()
	require.NoError(err)
	require.Equal(scanner.Position{1, 0}, s.Position())

	err = s.RequireKeyword("EXPLAIN")
	require.NoError(err)
	require.Equal(scanner.Position{1, 7}, s.Position())

	err = s.RequireWhitespace()
	require.NoError(err)
	require.Equal(scanner.Position{1, 8}, s.Position())

	err = s.RequireKeyword("ANALYZE")
	require.NoError(err)
	require.Equal(scanner.Position{1, 15}, s.Position())

	err = s.RequireWhitespace()
	require.NoError(err)
	require.Equal(scanner.Position{1, 16}, s.Position())

	err = s.RequireKeyword("SELECT")
	require.NoError(err)
	require.Equal(scanner.Position{2, 6}, s.Position())
}

func TestRequireKeyword(t *testing.T) {
	require := require.New(t)
	for _, tc := range []struct {
		Label string
		Input string
		Error string
	}{{
		Label: "empty",
		Input: "",
		Error: "keyword expected",
	}, {
		Label: "lowercase",
		Input: "select",
	}, {
		Label: "uppercase",
		Input: "SELECT",
	}, {
		Label: "partial",
		Input: "SEL",
		Error: "keyword expected",
	}, {
		Label: "extra",
		Input: "SELECTED",
	}} {
		s, err := scanner.New(tc.Input)
		require.NoError(err)

		err = s.RequireKeyword("SELECT")
		if tc.Error == "" {
			require.Nil(err, tc.Label)
		} else {
			require.Error(err, tc.Label)
			require.Equal(tc.Error, err.Error(), tc.Label)
		}
	}
}

func TestRequireWhitespace(t *testing.T) {
	require := require.New(t)
	for _, tc := range []struct {
		Label string
		Input string
		Error string
	}{{
		Label: "empty",
		Input: "",
		Error: "whitespace expected",
	}, {
		Label: "missing",
		Input: "foo",
		Error: "whitespace expected",
	}, {
		Label: "present",
		Input: " \r\n",
	}} {
		s, err := scanner.New(tc.Input)
		require.NoError(err)

		err = s.RequireWhitespace()
		if tc.Error == "" {
			require.Nil(err, tc.Label)
		} else {
			require.Error(err, tc.Label)
			require.Equal(tc.Error, err.Error(), tc.Label)
		}
	}
}

func TestSkipWhitespace(t *testing.T) {
	require := require.New(t)
	for _, tc := range []struct {
		Label string
		Input string
		Error string
	}{{
		Label: "empty",
		Input: "",
	}, {
		Label: "missing",
		Input: "foo",
	}, {
		Label: "present",
		Input: " \r\n",
	}} {
		s, err := scanner.New(tc.Input)
		require.NoError(err)

		err = s.SkipWhitespace()
		if tc.Error == "" {
			require.Nil(err, tc.Label)
		} else {
			require.Error(err, tc.Label)
			require.Equal(tc.Error, err.Error(), tc.Label)
		}
	}
}

func TestSnapshotAndReset(t *testing.T) {
	require := require.New(t)

	s, err := scanner.New("COMMIT")
	require.NoError(err)

	snapshot := s.Snapshot()

	err = s.RequireKeyword("COMMIT")
	require.NoError(err)

	err = s.RequireKeyword("COMMIT")
	require.Error(err)

	s.Reset(snapshot)

	err = s.RequireKeyword("COMMIT")
	require.NoError(err)
}
