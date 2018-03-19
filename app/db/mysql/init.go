package mysql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/baihym/go_eth_explorer/app/config"
)

var (
	__DB *sql.DB
)

func init() {
	var err error

	defer func() {
		if err := recover(); err != nil {
			DBClose()
			panic(err)
		}
	}()

	__DB, err = sql.Open("mysql", config.DBDSN)
	if err != nil {
		panic(err)
	}
	__DB.SetMaxOpenConns(config.DBMaxOpen)
	__DB.SetMaxIdleConns(config.DBMaxIdle)
	__DB.SetConnMaxLifetime(time.Second * 300)
	if err = __DB.Ping(); err != nil {
		panic(err)
	}
}

func DB() *sql.DB {
	return __DB
}

// DBClose 释放主库的资源
func DBClose() error {
	if __DB != nil {
		return __DB.Close()
	}
	return nil
}

func DBExec(query string, args ...interface{}) (sql.Result, error) {
	return __DB.Exec(query, args...)
}

func DBQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return __DB.Query(query, args...)
}

func DBQueryRow(query string, args ...interface{}) *sql.Row {
	return __DB.QueryRow(query, args...)
}

func DBExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return __DB.ExecContext(ctx, query, args...)
}

func DBQueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return __DB.QueryContext(ctx, query, args...)
}

func DBQueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return __DB.QueryRowContext(ctx, query, args...)
}
