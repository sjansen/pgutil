// +build docker,integration

package main_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

func connect() (c *pg.Conn, err error) {
	options := &pg.Options{
		Log: logger.Discard(),

		Database:       "pgutil_test_complete",
		ConnectRetries: 3,
		SSLMode:        "prefer",
	}
	return pg.New(options)
}

const actualPath = "testdata/actual.hcl"
const expectedPath = "testdata/expected.hcl"

func TestConnectAndQuery(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	expected, err := ioutil.ReadFile(expectedPath)
	require.NoError(err)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	db, err := c.InspectDatabase(&pg.InspectOptions{
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
