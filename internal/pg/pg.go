package pg

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/zapadapter"
	"go.uber.org/zap"
)

type Options struct {
	Log *zap.SugaredLogger

	Host     string
	Port     uint16
	Username string
	Password string
	Database string
}

type Conn struct {
	conn *pgx.Conn
	log  *zap.SugaredLogger
}

func New(o *Options) (*Conn, error) {
	cfg, err := pgx.ParseEnvLibpq()
	if err != nil {
		return nil, err
	}

	cfg = cfg.Merge(pgx.ConnConfig{
		Logger: zapadapter.NewLogger(
			o.Log.Desugar(),
		),

		Host:     o.Host,
		Port:     o.Port,
		User:     o.Username,
		Password: o.Password,
		Database: o.Database,
	})

	conn, err := pgx.Connect(cfg)
	if err != nil {
		return nil, err
	}

	c := &Conn{
		conn: conn,
		log:  o.Log,
	}
	return c, nil
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) Exec(query string) error {
	c.log.Debugw("executing query", "query", query)
	tag, err := c.conn.Exec(query)
	c.log.Debugf("rows affected = %d", tag.RowsAffected())
	return err
}

func (c *Conn) ServerVersion() (string, error) {
	c.log.Infow("requesting server version")

	var version string
	c.log.Debugw("executing query", "query", "SELECT VERSION()")
	err := c.conn.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	c.log.Debugf("server version = %q", version)
	return version, nil
}
