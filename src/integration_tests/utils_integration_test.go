package integration_tests

import (
	"testing"

	"golang-kafka-postgre/src/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseIntegration(t *testing.T) {
	err := utils.InitDB("user=postgres password=yourRoamer14 host=localhost port=5432 dbname=gokafka-db sslmode=disable")
	require.NoError(t, err, "should connect to the database")

	// Test adding data to the table log
	utils.AddDataToTableLog("error test", 500, "{\"key\":\"value\"}", "koli123", "key1", "test_table")

	count, err := utils.CountEntries("test_table", "key1")
	require.NoError(t, err, "should query the database successfully")
	assert.Equal(t, 1, count, "should find one message in the database")
}
