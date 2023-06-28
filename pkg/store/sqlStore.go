package store

import "database/sql"

type sqlStore struct {
	DB *sql.DB
}
func NewSqlStore(database *sql.DB) StoreInterface {
	return &sqlStore{
		DB: database,
	}
}
