package mocks

import "github.com/stretchr/testify/mock"

type DB struct {
	mock.Mock
}

func (db *DB) Close() error {
	result := db.Called()
	return result.Error(0)
}

func (db *DB) Exec(query string) error {
	result := db.Called(query)
	return result.Error(0)
}

func (db *DB) ServerVersion() (string, error) {
	result := db.Called()
	return result.String(0), result.Error(1)
}
