package testutil

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 33306
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}
	db, err := sqlx.Open(
		"mysql",
		fmt.Sprintf("todo:todo@tcp(localhost:%d)/todo_test?parseTime=true", port),
	)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	t.Cleanup(
		func() {
			_ := db.Close()
		},
	)
	return sqlx.NewDb(db.DB, "mysql")
}
