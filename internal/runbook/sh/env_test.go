package sh_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/sh"
)

func TestEnv(t *testing.T) {
	require := require.New(t)

	environ := []string{
		"USER=sjansen",
		"HOME=/home/sjansen",
	}

	// Zero Value
	e := &sh.Env{}
	actual := e.Apply(environ)
	require.Equal(environ, actual, "Zero Value")

	// Allow
	e = &sh.Env{Allow: []string{"USER"}}
	expected := []string{
		"USER=sjansen",
	}
	actual = e.Apply(environ)
	require.Equal(expected, actual, "Allow")

	// Deny
	e = &sh.Env{Deny: []string{"USER"}}
	expected = []string{
		"HOME=/home/sjansen",
	}
	actual = e.Apply(environ)
	require.Equal(expected, actual, "Deny")

	// Set
	e = &sh.Env{Set: map[string]string{
		"USER": "root",
		"HOME": "/root",
	}}
	expected = []string{
		"USER=root",
		"HOME=/root",
	}
	actual = e.Apply(environ)
	require.Equal(expected, actual, "Set")

	// Allow & Set
	e = &sh.Env{
		Allow: []string{"HOME"},
		Set: map[string]string{
			"FOO": "bar",
		},
	}
	expected = []string{
		"HOME=/home/sjansen",
		"FOO=bar",
	}
	actual = e.Apply(environ)
	require.Equal(expected, actual, "Allow & Set")

	// Deny & Set
	e = &sh.Env{
		Deny: []string{"USER"},
		Set: map[string]string{
			"FOO": "baz",
		},
	}
	expected = []string{
		"HOME=/home/sjansen",
		"FOO=baz",
	}
	actual = e.Apply(environ)
	require.Equal(expected, actual, "Deny & Set")
}
