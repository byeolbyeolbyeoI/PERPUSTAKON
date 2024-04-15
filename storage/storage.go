package storage

import (
	"github.com/gofiber/storage/mysql/v2"

	"fmt"
	"os"

	"perpustakaan/config"
)

func createStorage() *mysql.Storage {
	config := config.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Protocol: os.Getenv("DB_PROTOCOL"),
		Path:     os.Getenv("DB_PATH"),
		DBName:   os.Getenv("DB_DBNAME"),
	}

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", config.User, config.Password, config.Protocol, config.Path, config.DBName)

	store := mysql.New(mysql.Config{
		ConnectionURI: dsn,
	})

	return store
}
