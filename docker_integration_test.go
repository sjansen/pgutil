//go:build docker && integration
// +build docker,integration

package main_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

func connect(ctx context.Context) (c *pg.Conn, err error) {
	options := &pg.Options{
		Log: logger.Discard(),

		Database:       "pgutil_test_complete",
		ConnectRetries: 3,
		SSLMode:        "prefer",
	}
	return pg.New(ctx, options)
}

const actualPath = "testdata/actual.hcl"
const expectedPath = "testdata/expected.hcl"

func TestConnectAndQuery(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	expected, err := ioutil.ReadFile(expectedPath)
	require.NoError(err)

	ctx := context.TODO()
	c, err := connect(ctx)
	require.NoError(err)
	defer c.Close(ctx)

	db, err := c.InspectDatabase(ctx, &pg.InspectOptions{
		SortColumns: true,
		SortIndexes: true,
	})
	require.NoError(err)

	buf := &bytes.Buffer{}
	err = db.Write(buf)
	require.NoError(err)

	actual := buf.Bytes()
	if !assert.Equal(string(expected), string(actual)) {
		ioutil.WriteFile(actualPath, actual, 0666)
		t.Log(
			"Temp JSON file created to facilitate debugging.",
			"\nexpected:", expectedPath,
			"\nactual:", actualPath,
		)
	}
}
