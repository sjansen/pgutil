package pg

import (
	"os"
	"os/user"

	gopg "github.com/go-pg/pg"
	"go.uber.org/zap"
)

type Options struct {
	Log *zap.SugaredLogger

	Network  string
	Address  string
	Username string
	Password string
	Database string
}

type Pool struct {
	conn *gopg.DB
	log  *zap.SugaredLogger
}

func New(o *Options) (*Pool, error) {
	options, err := mergeOptions(o)
	if err != nil {
		return nil, err
	}

	p := &Pool{
		conn: gopg.Connect(
			options,
		),
		log: o.Log,
	}
	return p, nil
}

func mergeOptions(o *Options) (*gopg.Options, error) {
	result := &gopg.Options{
		ApplicationName: "pgutil",
	}

	if o.Address != "" {
		result.Addr = o.Address
	} else if host := os.Getenv("PGHOST"); host != "" {
		if port := os.Getenv("PGPORT"); port != "" {
			result.Addr = host + ":" + port
		} else {
			result.Addr = host + ":5432"
		}
	}

	if o.Username != "" {
		result.User = o.Username
	} else if username := os.Getenv("PGUSER"); username != "" {
		result.User = username
	} else if u, err := user.Current(); err != nil {
		return nil, err
	} else {
		result.User = u.Username
	}

	if o.Password != "" {
		result.Password = o.Password
	} else if password := os.Getenv("PGPASSWORD"); password != "" {
		result.Password = password
	}

	if o.Database != "" {
		result.Database = o.Database
	} else if database := os.Getenv("PGDATABASE"); database != "" {
		result.Database = database
	}

	return result, nil
}

func (p *Pool) Close() error {
	return p.conn.Close()
}

func (p *Pool) Exec(query string) error {
	p.log.Debugw("executing query", "query", query)
	_, err := p.conn.Exec(query)
	return err
}

func (p *Pool) ServerVersion() (string, error) {
	p.log.Infow("requesting server version")

	var version string
	p.log.Debugw("executing query", "query", "SELECT VERSION()")
	_, err := p.conn.Query(gopg.Scan(&version), "SELECT VERSION()")
	if err != nil {
		return "", err
	}

	p.log.Debugf("server version = %q", version)
	return version, nil
}
