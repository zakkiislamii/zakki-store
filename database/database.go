package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
)

var DbConnection *sql.DB

func DbMigrate(dbParam *sql.DB) error {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./sql_migrations"),
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	DbConnection = dbParam
	fmt.Println("Applied", n, "migrations!")
	return nil
}
