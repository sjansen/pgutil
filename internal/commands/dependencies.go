package commands

type Dependencies struct {
	DB func(map[string]string) (DB, error)
}

type DB interface {
	Close() error
	ServerVersion() (string, error)
}
