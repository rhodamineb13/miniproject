package main

import (
	"log"
	"miniproject/common/database"
)

func main() {
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
}
