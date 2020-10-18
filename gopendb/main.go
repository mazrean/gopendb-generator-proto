package gopendb

import (
	"database/sql"

	"github.com/mazrean/gopendb-generator-proto/gopendb/users"
)

type Users = users.Nullable
type UsersTable = users.Table

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) Users() *users.Chain {
	return users.NewChain(db.db)
}
