package db

import (
	"database/sql"
	"os"
)

func RunMigrations(db *sql.DB) error {
	migrationFiles := []string{
		"migrations/000005_create_contact_table.up.sql",
	}

	for _, file := range migrationFiles {
		migrationSQL, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(migrationSQL))
		if err != nil {
			return err
		}
	}

	return nil
}
