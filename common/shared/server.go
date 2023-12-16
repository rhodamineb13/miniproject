package shared

import (
	"log"
	"miniproject/common/database"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
	r  *redis.Client
}

var DB = NewDatabase()

func NewDatabase() *Database {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := database.NewRedisClient()

	return &Database{
		db: db,
		r:  r,
	}
}
