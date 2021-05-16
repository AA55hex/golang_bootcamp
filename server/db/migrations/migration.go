package migrations

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/mysql"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB) error {
	fmt.Println("Starting migrations")
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return errors.New("mysql.WithInstance: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations/", "mysql", driver)
	if err != nil {
		return errors.New("migrate.NewWithDatabaseInstance: " + err.Error())
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.New("m.Steps: " + err.Error())
	}
	fmt.Println("Migrations executed")
	return nil
}
