package config

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@unix(/run/mysqld/mysqld.sock)/perpustakaan")
	if err != nil {
		return nil, err
	}

	return db, nil
}
