package config

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type Config struct {
	User string
	Password string
	Protocol string
	Path string
	DBName string
}

func Connect(config *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", config.User, config.Password, config.Protocol, config.Path, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
