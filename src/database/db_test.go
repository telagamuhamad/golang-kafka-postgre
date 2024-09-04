package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	InitDB()
	defer DB.Close()

	err := DB.Ping()
	assert.NoError(t, err, "should connect to the database")
}
