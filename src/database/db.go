package database

import (
	"database/sql"
	"fmt"
	"log"

	"golang-kafka-postgre/src/constant"

	_ "github.com/lib/pq" // Import Postgres driver
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", constant.DatabaseUser, constant.DatabasePassword, constant.DatabaseHost)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
}
