package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/symball/go-gin-boilerplate/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/schema"
)

var dbHandle *bun.DB

func DBGet() *bun.DB {
	return dbHandle
}

func DBInit(ctx context.Context) error {
	driver := strings.ToLower(strings.TrimSpace(
		config.AppConfig.DBDriver,
	))
	dsn := config.AppConfig.DBDSN
	isVerbose := config.AppConfig.Verbose

	if driver == "" {
		driver = detectDriver(dsn)
	}

	log.Printf("Database driver (%s) with DSB: %s", driver, dsn)

	var (
		sqlDB   *sql.DB
		dialect schema.Dialect
		err     error
	)

	switch driver {
	case "sqlite", "sqlite3":
		sqlDB, err = sql.Open(sqliteshim.ShimName, dsn)
		if err != nil {
			return fmt.Errorf("open sqlite: %w", err)
		}

		dialect = sqlitedialect.New()

	case "postgres", "postgresql":
		sqlDB = sql.OpenDB(
			pgdriver.NewConnector(
				pgdriver.WithDSN(dsn),
			),
		)
		dialect = pgdialect.New()

	case "mysql", "mariadb":
		sqlDB, err = sql.Open("mysql", dsn)
		if err != nil {
			return fmt.Errorf("open mysql: %w", err)
		}

		dialect = mysqldialect.New()

	case "mssql", "sqlserver":
		sqlDB, err = sql.Open("sqlserver", dsn)
		if err != nil {
			return fmt.Errorf("open sql server: %w", err)
		}

		dialect = mssqldialect.New()

	case "oracle":
		return fmt.Errorf(
			"oracle is not supported by Bun's built-in dialects",
		)

	default:
		return fmt.Errorf("database driver not supported: %q", driver)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return fmt.Errorf("ping %s database: %w", driver, err)
	}

	dbHandle = bun.NewDB(sqlDB, dialect)
	dbHandle.WithQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(isVerbose),
	))

	return nil
}

func detectDriver(dsn string) string {
	dsn = strings.ToLower(strings.TrimSpace(dsn))

	switch {
	case strings.HasPrefix(dsn, "file:"),
		strings.HasPrefix(dsn, ":memory:"):
		return "sqlite"

	case strings.HasPrefix(dsn, "postgres://"),
		strings.HasPrefix(dsn, "postgresql://"):
		return "postgres"

	case strings.HasPrefix(dsn, "mysql://"),
		strings.Contains(dsn, "@tcp("):
		return "mysql"

	case strings.HasPrefix(dsn, "sqlserver://"),
		strings.HasPrefix(dsn, "mssql://"):
		return "mssql"

	default:
		return ""
	}
}
