// +build integration

package pg_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

func connect() (p *pg.Pool, err error) {
	options := &pg.Options{
		Log: logger.Discard(),
	}
	for retries := 0; retries < 5; retries++ {
		if p, err = pg.New(options); err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
}

func TestConnectAndQuery(t *testing.T) {
	require := require.New(t)

	p, err := connect()
	require.NoError(err)
	defer p.Close()

	version, err := p.ServerVersion()
	require.NoError(err)
	require.NotEmpty(version)
}

func TestExec(t *testing.T) {
	require := require.New(t)

	p, err := connect()
	require.NoError(err)
	defer p.Close()

	query := `
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
`
	err = p.Exec(query)
	require.NoError(err)
}
