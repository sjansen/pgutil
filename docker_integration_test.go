// +build docker,integration

package main_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/ddl"
	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

func connect() (c *pg.Conn, err error) {
	options := &pg.Options{
		Log: logger.Discard(),

		ConnectRetries: 3,
		SSLMode:        "prefer",
	}
	return pg.New(options)
}

func TestConnectAndQuery(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	expected, err := ioutil.ReadFile("testdata/expected.hcl")
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
	err = ddl.Write(buf, db)
	require.NoError(err)

	actual := buf.Bytes()
	if !assert.Equal(string(expected), string(actual)) {
		ioutil.WriteFile("testdata/actual.hcl", actual, 0666)
	}
}
