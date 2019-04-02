package loader_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/loader"
)

func TestErrors(t *testing.T) {
	require := require.New(t)

	var err error = &loader.InternalError{Original: errors.New("i am error")}
	require.NotEmpty(err.Error())
}
