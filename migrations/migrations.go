package migrations

import (
	"fmt"

	"github.com/symball/go-gin-boilerplate/migrations/mssql"
	"github.com/symball/go-gin-boilerplate/migrations/mysql"
	"github.com/symball/go-gin-boilerplate/migrations/postgres"
	"github.com/symball/go-gin-boilerplate/migrations/sqlite"
	"github.com/symball/go-gin-boilerplate/storage"
	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/migrate"
)

func GetForDatabase() (*migrate.Migrations, error) {
	db := storage.DBGet()

	if db == nil {
		return nil, fmt.Errorf("database has not been initialized")
	}

	switch db.Dialect().Name() {
	case dialect.PG:
		return postgres.Migrations, nil

	case dialect.MySQL:
		return mysql.Migrations, nil

	case dialect.MSSQL:
		return mssql.Migrations, nil

	case dialect.SQLite:
		return sqlite.Migrations, nil

	default:
		return nil, fmt.Errorf(
			"no migrations available for database dialect %q",
			db.Dialect().Name(),
		)
	}
}
