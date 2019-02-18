// +build integration

package db_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/db"
)

func connect() (c *db.DB, err error) {
	options := map[string]string{}
	for retries := 0; retries < 5; retries++ {
		if c, err = db.New(options); err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
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
