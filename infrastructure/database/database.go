package database

import (
	"context"
	"database/sql"
	"time"
)

var postgres = "postgres"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string

	MaxOpenConn int
	MaxIdleConn int

	MaxLifeTime time.Duration
	MaxIdleTime time.Duration
}

type Database struct {
	*sql.DB
}

func (db *Database) Query(ctx context.Context, query string, any ...interface{}) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, any...)
}

func (db *Database) Exec(ctx context.Context, query string, any ...interface{}) (sql.Result, error) {
	return db.ExecContext(ctx, query, any...)
}

func (db *Database) QueryRow(ctx context.Context, query string, any ...interface{}) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, any...)
}
