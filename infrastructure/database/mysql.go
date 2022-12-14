package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(config Config) *Database {
	if config.SSLMode != "" {
		config.SSLMode = fmt.Sprintf("?sslmode=%s", config.SSLMode)
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName)

	db, err := sql.Open(mysql, connStr)
	if err != nil {
		log.Fatal("could not open mysql database", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("could not connect mysql database", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)

	db.SetConnMaxLifetime(config.MaxLifeTime)
	db.SetConnMaxIdleTime(config.MaxIdleTime)

	return &Database{DB: db}
}
