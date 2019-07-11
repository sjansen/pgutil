// +build integration

package pg_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

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
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	version, err := c.ServerVersion()
	require.NoError(err)
	require.NotEmpty(version)
}

func TestDescribeFunction(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	function, err := c.DescribeFunction("pg_catalog", "current_database")
	require.NoError(err)
	require.NotNil(function)

	function, err = c.DescribeFunction("public", "no_such_function")
	require.Equal(pg.ErrNotFound, err)
	require.Nil(function)
}

func TestDescribeTrigger(t *testing.T) {
	if os.Getenv("PGUTIL_TEST_TAGS") == "" {
		t.Skip("missing $PGUTIL_TEST_TAGS")
	}
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	trigger, err := c.DescribeTrigger("public", "measurement", "update_modified_column")
	require.NoError(err)
	require.NotNil(trigger)

	trigger, err = c.DescribeTrigger("public", "no_such_table", "no_such_trigger")
	require.Equal(pg.ErrNotFound, err)
	require.Nil(trigger)
}

func TestExec(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	query := `
BEGIN
;
CREATE TABLE IF NOT EXISTS measurements (
    id BIGSERIAL NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    value DOUBLE PRECISION NOT NULL
)
;
DELETE FROM measurements
WHERE timestamp < now() - interval '5 minutes'
;
INSERT INTO measurements
    (timestamp, value)
VALUES
    (now(), random())
;
INSERT INTO measurements
    (timestamp, value)
VALUES
    (now(), random())
;
ROLLBACK
;
`
	err = c.Exec(query)
	require.NoError(err)
}

func TestListColumns(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	columns, err := c.ListColumns("pg_catalog", "pg_class")
	require.NoError(err)
	require.NotEmpty(columns)
}

func TestListFunctions(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	functions, err := c.ListFunctions()
	require.NoError(err)
	require.NotEmpty(functions)
}

func TestListSchemas(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	schemas, err := c.ListSchemas()
	require.NoError(err)
	require.NotEmpty(schemas)
}

func TestListTables(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	tables, err := c.ListTables()
	require.NoError(err)
	require.NotEmpty(tables)
}
