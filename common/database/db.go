package database

import (
	"fmt"

	"miniproject/common/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBname)
	db, err := sqlx.Open("pgx", conn)
	if err != nil {
		return nil, err
	}

	err = MigrateUp(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migration", "testproject", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
