package database

import (
	"fmt"
	"miniproject/common/config"

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

	return db, nil
}
