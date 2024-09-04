package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitDB(t *testing.T) {
	// Mock database connection string
	dataSourceName := "user=postgres password=yourRoamer14 host=localhost port=5432 dbname=gokafka-db sslmode=disable"

	// Initialize the database
	err := InitDB(dataSourceName)
	require.NoError(t, err, "should connect to the database")

	// Check if the database connection is established
	err = db.Ping()
	assert.NoError(t, err, "should connect to the database successfully")
}

func TestAddDataToTableLog(t *testing.T) {
	// Mock database connection string
	dataSourceName := "user=postgres password=yourRoamer14 host=localhost port=5432 dbname=gokafka-db sslmode=disable"
	err := InitDB(dataSourceName)
	require.NoError(t, err, "should connect to the database")

	// Use an in-memory SQLite database for testing purposes
	// Replace this with actual database setup if necessary
	db.Exec("CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, error_message TEXT, status_code INT, json_data TEXT, koli_number TEXT, key1 TEXT, created_at TIMESTAMP)")

	// Insert test data
	AddDataToTableLog("Test Error", 500, `{"key":"value"}`, "123", "testKey", "test_table")

	// Verify data is inserted
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM test_table WHERE error_message = $1", "Test Error").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "should insert log data into the table")
}

func TestPushToCore(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	data := []byte(`{"key":"value"}`)
	resp, err := PushToCore(data, server.URL, http.MethodPost)
	if err != nil {
		t.Fatalf("PushToCore() error = %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
