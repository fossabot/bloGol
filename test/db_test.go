package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bloGol/bloGol/internal/database"
	"github.com/stretchr/testify/assert"
)

var testDB *database.DataBase

func TestDataBaseOpenInvalid(t *testing.T) {
	var err error
	testDB, err = database.Open("/wtf/dude")
	assert.Error(t, err)
}

func TestDataBaseOpenValid(t *testing.T) {
	var validPath = filepath.Join(
		os.Getenv("GOPATH"), "src", "github.com", "bloGol", "bloGol", "test", "development.db",
	)

	var err error
	testDB, err = database.Open(validPath)
	assert.NoError(t, err)
}
