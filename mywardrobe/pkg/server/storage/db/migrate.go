package postgresql

import (
	"database/sql"

	"example.com/m/v2/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func build(db *sql.DB, migrationFile string) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.StdLogger().Fatal(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationFile,
		"postgres", driver)
	m.Up()
}
