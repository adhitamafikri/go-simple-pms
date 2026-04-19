package repositories

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}
