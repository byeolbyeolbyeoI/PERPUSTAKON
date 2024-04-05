package repository

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}
