// Package contains command migrations
package migrations

import (
	"database/sql"

	"test-be-IMP/pkg/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	dbs, _    = sql.Open(db.Get().Driver, db.Get().DSN())
	driver, _ = mysql.WithInstance(dbs, &mysql.Config{})
)

func Up() error {
	m, err := migrate.NewWithDatabaseInstance(
		cfg.Migration.Source,
		db.Get().Driver,
		driver,
	)
	if nil != err {
		return err
	}
	err = m.Up()

	return ignoreNoChange(err)
}

func Down() error {
	m, err := migrate.NewWithDatabaseInstance(
		cfg.Migration.Source,
		db.Get().Driver,
		driver,
	)
	if nil != err {
		return err
	}
	err = m.Down()

	return ignoreNoChange(err)
}

func Drop() error {
	m, err := migrate.NewWithDatabaseInstance(
		cfg.Migration.Source,
		db.Get().Driver,
		driver,
	)
	if nil != err {
		return err
	}
	err = m.Drop()

	return ignoreNoChange(err)
}

func ignoreNoChange(err error) error {
	if nil != err && err == migrate.ErrNoChange {
		return nil
	}

	return err
}
