package testutil

import (
	"context"
	"os"
	"strings"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

func Connect(ctx context.Context) (c *pg.Conn, err error) {
	options := &pg.Options{
		Log: logger.Discard(),

		Host:           "pg9.6",
		Database:       "pgutil_test_complete",
		ConnectRetries: 3,
		SSLMode:        "prefer",
	}
	return pg.New(ctx, options)
}

func PGHosts() []string {
	hosts := strings.Split(os.Getenv("PGUTIL_TEST_HOSTS"), " ")
	if len(hosts) < 1 {
		hosts = []string{""}
	}
	return hosts
}
