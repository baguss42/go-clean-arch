package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgres(config Config) *Database {
	if config.SSLMode != "" {
		config.SSLMode = fmt.Sprintf("?sslmode=%s", config.SSLMode)
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s%s",
		config.Username,
		config.Password,
		config.Host,
		config.DBName,
		config.SSLMode)

	db, err := sql.Open(postgres, connStr)
	if err != nil {
		log.Fatal("could not open postgres database", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("could not connect postgres database", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)

	db.SetConnMaxLifetime(config.MaxLifeTime)
	db.SetConnMaxIdleTime(config.MaxIdleTime)

	return &Database{DB: db}
}
