package _interface

import (
	"context"
	"database/sql"
)

type DatabaseInterface interface {
	Exec(ctx context.Context, query string, any ...interface{}) (sql.Result, error)
	Query(ctx context.Context, query string, any ...interface{}) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, any ...interface{}) (*sql.Rows, error)
}
