package storage

import (
	"database/sql"
	"log"
	"strings"

	"github.com/symball/go-gin-boilerplate/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

var dbHandle *bun.DB

// Initiate an SQLite DB instance
func DBInit() {
	dsn := config.AppConfig.DBDSN
	switch true {
	// SQLite
	case strings.HasPrefix(dsn, "file"):
		sqldb, sqliteError := sql.Open(sqliteshim.ShimName, dsn)
		if sqliteError != nil {
			log.Fatalf("failed to open sqlite: %v", sqliteError)
		}
		dbHandle = bun.NewDB(sqldb, sqlitedialect.New())
	case strings.HasPrefix(dsn, "postgres"):

		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		dbHandle = bun.NewDB(sqldb, pgdialect.New())
	default:
		log.Fatalf("database not supported: %s", dsn)
	}

	dbHandle.AddQueryHook(bundebug.NewQueryHook())
}

// Retrieve the DB instance
func DBGet() *bun.DB {
	return dbHandle
}
