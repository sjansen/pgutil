package dbmock

import "github.com/stretchr/testify/mock"

type DB struct {
	mock.Mock
}

func (db *DB) Close() error {
	args := db.Called()
	return args.Error(0)
}

func (db *DB) ServerVersion() (string, error) {
	args := db.Called()
	return args.String(0), args.Error(1)
}
