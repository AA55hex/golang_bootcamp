package connection

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// trying to execute migration scripts
// todo: add .env support
func try_migrate(db *sql.DB) error {

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
	switch err = m.Up(); err {
	case nil:
		fmt.Println("Migrations executed")
	case migrate.ErrNoChange:
		fmt.Println("No migrations to execute")
	default:
		return errors.New("m.Steps: " + err.Error())
	}

	return nil
}
