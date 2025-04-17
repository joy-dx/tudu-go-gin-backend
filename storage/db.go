package storage

import (
	"database/sql"
	"github.com/symball/go-gin-boilerplate/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var dbHandle *bun.DB

// Initiate an SQLite DB instance
func DBInit() {
	dsn := config.AppConfig.PostgresDSN
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	dbHandle = bun.NewDB(sqldb, pgdialect.New())
	dbHandle.AddQueryHook(bundebug.NewQueryHook())
}

// Retrieve the DB instance
func DBGet() *bun.DB {
	return dbHandle
}
