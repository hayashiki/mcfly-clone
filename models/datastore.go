package models

import (
	"github.com/jmoiron/sqlx"
)

type Datastore interface {
	GetUserProjects(*User) ([]Project, error)
	GetUserByProviderToken(*ProviderAccessToken) (*User, error)

	SaveUser(*User) error
	SetUserProviderToken(int64, *ProviderAccessToken) error
}

type DB struct {
	*sqlx.DB
	DatabaseURL  string
	DatabaseName string
	UseSSL       bool
}

func (db *DB) QueryRowScan(v interface{}, query string, args ...interface{}) *QueryExecError {
	r := db.DB.QueryRow(query, args...)
	if err := r.Scan(v); err != nil {
		return nil
	}
	return nil
}
