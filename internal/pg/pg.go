package pg

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/zapadapter"
	"go.uber.org/zap"
)

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

type Conn struct {
	conn *pgx.Conn
	log  *zap.SugaredLogger
}

var ErrNoHostForTLS = errors.New("host server name must be provided when TLS is required")

func New(o *Options) (*Conn, error) {
	cfg, err := pgx.ParseEnvLibpq()
	if err != nil {
		return nil, err
	}

	cfg.Host = firstOfString(o.Host, cfg.Host)
	cfg.Port = firstOfUint16(o.Port, cfg.Port)
	cfg.User = firstOfString(o.Username, cfg.User)
	cfg.Password = firstOfString(o.Password, cfg.Password)
	cfg.Database = firstOfString(o.Database, cfg.Database)

	sslmode := o.SSLMode
	if sslmode == "" && cfg.Password != "" {
		sslmode = "verify-full"
	}
	if sslmode != "" {
		err = applySSLMode(&cfg, sslmode)
		if err != nil {
			return nil, err
		}
	}

	cfg.Logger = zapadapter.NewLogger(
		o.Log.Desugar(),
	)

	delay := 100 * time.Millisecond
	var conn *pgx.Conn
	for retries := 0; retries <= o.ConnectRetries; retries++ {
		conn, err = pgx.Connect(cfg)
		if err != nil {
			if retries < o.ConnectRetries {
				time.Sleep(delay)
				delay = time.Duration(2^retries) * 100 * time.Millisecond
			} else {
				return nil, err
			}
		}
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
	if err != nil {
		c.log.Debugf("rows affected = %d", tag.RowsAffected())
	}
	return err
}

func applySSLMode(cfg *pgx.ConnConfig, sslmode string) error {
	switch sslmode {
	case "disable":
		cfg.UseFallbackTLS = false
		cfg.TLSConfig = nil
		cfg.FallbackTLSConfig = nil
	case "allow":
		cfg.UseFallbackTLS = true
		cfg.TLSConfig = nil
		cfg.FallbackTLSConfig = &tls.Config{InsecureSkipVerify: true}
	case "prefer":
		cfg.UseFallbackTLS = true
		cfg.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		cfg.FallbackTLSConfig = nil
	case "require":
		cfg.UseFallbackTLS = false
		cfg.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		cfg.FallbackTLSConfig = nil
	case "verify-full":
		cfg.UseFallbackTLS = false
		cfg.TLSConfig = &tls.Config{ServerName: cfg.Host}
		cfg.FallbackTLSConfig = nil
	default:
		return errors.New("invalid sslmode")
	}
	return nil
}

func firstOfString(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func firstOfUint16(a, b uint16) uint16 {
	if a != 0 {
		return a
	}
	return b
}
