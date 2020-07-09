package seeder

import (
	"database/sql"

	"github.com/go-testfixtures/testfixtures/v3"
)

// Seed the database
func Seed(db *sql.DB) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),                 // You database connection
		testfixtures.Dialect("sqlite"),            // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("seeder/fixtures"), // the directory containing the YAML files
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	if err != nil {
		return err
	}
	return fixtures.Load()
}
