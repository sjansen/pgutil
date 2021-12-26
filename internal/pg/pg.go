package pg

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"go.uber.org/zap"
)

// Options contains settings for connecting to a PostgreSQL database
type Options struct {
	Log *zap.SugaredLogger

	Host     string
	Port     uint16
	SSLMode  string
	Username string
	Password string
	Database string

	ConnectRetries int
}

func (o *Options) connstring() string {
	parts := make([]string, 0, 6)
	if o.Host != "" {
		parts = append(parts, "host="+o.Host)
	}
	if o.Port != 0 {
		parts = append(parts, "port="+strconv.Itoa(int(o.Port)))
	}
	if o.SSLMode != "" {
		parts = append(parts, "sslmode="+o.SSLMode)
	}
	if o.Username != "" {
		parts = append(parts, "user="+o.Username)
	}
	if o.Password != "" {
		parts = append(parts, "password="+o.Password)
	}
	if o.Database != "" {
		parts = append(parts, "dbname="+o.Database)
	}
	return strings.Join(parts, " ")
}

// Conn is a connection to a PostgreSQL database
type Conn struct {
	Host     string
	Database string

	conn *pgx.Conn
	log  *zap.SugaredLogger
}

// New connects to a PostgreSQL database
func New(ctx context.Context, o *Options) (*Conn, error) {
	cfg, err := pgx.ParseConfig(o.connstring())
	if err != nil {
		return nil, err
	}

	cfg.Logger = zapadapter.NewLogger(
		o.Log.Desugar(),
	)

	delay := 100 * time.Millisecond
	var conn *pgx.Conn
	for retries := 0; retries <= o.ConnectRetries; retries++ {
		conn, err = pgx.ConnectConfig(ctx, cfg)
		if err != nil {
			if retries < o.ConnectRetries {
				time.Sleep(delay)
				delay = time.Duration(2^retries) * 100 * time.Millisecond
			} else {
				return nil, err
			}
		}
	}

	host := cfg.Host
	if strings.Contains(host, string(os.PathSeparator)) {
		if host, err = os.Hostname(); err != nil {
			return nil, err
		}
	}

	database := cfg.Database
	if database == "" {
		database = cfg.User
	}

	c := &Conn{
		Host:     host,
		Database: database,
		conn:     conn,
		log:      o.Log,
	}
	return c, nil
}

// Close closes a connection
func (c *Conn) Close(ctx context.Context) error {
	return c.conn.Close(ctx)
}

// Exec executes SQL statements
func (c *Conn) Exec(ctx context.Context, sql string) error {
	c.log.Debugw("executing sql", "sql", sql)
	tag, err := c.conn.Exec(ctx, sql)
	if err != nil {
		c.log.Debugf("rows affected = %d", tag.RowsAffected())
	}
	return err
}
