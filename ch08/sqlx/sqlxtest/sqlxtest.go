package sqlxtest

import (
	"context"
	"testing"

	"github.com/inancgumus/effective-go/ch08/sqlx"
)

// DefaultTestDSN is the default test database DSN.
const DefaultTestDSN = ":memory:"

// Dial opens a test database connection.
func Dial(tb testing.TB) *sqlx.DB {
	tb.Helper()

	db, err := sqlx.Dial(context.Background(), sqlx.DefaultDriver, DefaultTestDSN)
	if err != nil {
		tb.Fatalf("dialing test db: %v", err)
	}
	tb.Cleanup(func() {
		if err := db.Close(); err != nil {
			tb.Log("closing test db:", err)
		}
	})
	return db
}

// You can add another Dial here for opening a non-memory database.
