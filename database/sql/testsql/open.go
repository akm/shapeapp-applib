package testsql

import (
	"database/sql"
	"os"
	"testing"

	"applib/log/slog"

	sqldbloggerslog "github.com/akm/sqldb-logger-slog"
	_ "github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"
)

var (
	dsn = os.Getenv("DB_DSN")
)

func Open(t *testing.T, logger *slog.Logger) *sql.DB {
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("unable to open database: %v", err)
	}

	adapter := sqldbloggerslog.New(logger)
	pool = sqldblogger.OpenDriver(dsn, pool.Driver(), adapter)

	return pool
}
