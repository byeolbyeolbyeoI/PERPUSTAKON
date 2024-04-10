package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Config struct {
	User     string
	Password string
	Protocol string
	Path     string
	DBName   string
}

func Connect() (*sql.DB, error) {
	config := Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Protocol: os.Getenv("DB_PROTOCOL"),
		Path:     os.Getenv("DB_PATH"),
		DBName:   os.Getenv("DB_DBNAME"),
	}

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", config.User, config.Password, config.Protocol, config.Path, config.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
