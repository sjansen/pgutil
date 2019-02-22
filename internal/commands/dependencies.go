package commands

import "io"

type Dependencies struct {
	DB      func(map[string]string) (DB, error)
	Process func([]string) Process
}

type Process interface {
	Run(stdout, stderr io.Writer) error
}

type DB interface {
	Close() error
	Exec(string) error
	ServerVersion() (string, error)
}
