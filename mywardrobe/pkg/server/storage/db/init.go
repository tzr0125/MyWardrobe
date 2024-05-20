package postgresql

import (
	"database/sql"

	"example.com/m/v2/pkg/logger"
	_ "github.com/lib/pq"
)

func InitPostgres(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.StdLogger().Fatal(err)
	}
	return db
}
