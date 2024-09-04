// src/utils/utils.go
package utils

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func AddDataToTableLog(errorMessage string, statusCode int, jsonData string, koliNumber string, key1 string, tableName string) {
	if db == nil {
		log.Println("Database connection is not initialized")
		return
	}

	query := `INSERT INTO ` + tableName + ` (error_message, status_code, json_data, koli_number, key1, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query, errorMessage, statusCode, jsonData, koliNumber, key1, time.Now())
	if err != nil {
		log.Printf("Error inserting log data: %v", err)
	}
}

func PushToCore(data []byte, url, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return client.Do(req)
}

// src/utils/utils.go
func CountEntries(tableName string, key1 string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM ` + tableName + ` WHERE key1 = $1`
	err := db.QueryRow(query, key1).Scan(&count)
	return count, err
}
