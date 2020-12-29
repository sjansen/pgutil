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
		switch tc.Error {
		case "":
			require.Nil(err)
		default:
			require.Equal(tc.Error, err.Error())
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
		Label: "invalid",
		Input: "foo",
		Error: "whitespace expected",
	}, {
		Label: "valid",
		Input: " foo ",
		Error: "",
	}} {
		s, err := scanner.New(tc.Input)
		require.NoError(err)

		err = s.RequireWhitespace()
		switch tc.Error {
		case "":
			require.Nil(err)
		default:
			require.Equal(tc.Error, err.Error())
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
		Error: "",
	}, {
		Label: "present",
		Input: " foo ",
	}, {
		Label: "missing",
		Input: "foo",
	}} {
		s, err := scanner.New(tc.Input)
		require.NoError(err)

		err = s.SkipWhitespace()
		switch tc.Error {
		case "":
			require.Nil(err)
		default:
			require.Equal(tc.Error, err.Error())
		}
	}
}
