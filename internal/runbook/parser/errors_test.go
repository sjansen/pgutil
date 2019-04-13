package parser_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
)

func TestErrors(t *testing.T) {
	require := require.New(t)

	var err error = &parser.InternalError{Original: errors.New("i am error")}
	require.NotEmpty(err.Error())
}
