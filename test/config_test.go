package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/stretchr/testify/assert"
)

var testConfig *config.Configuration

func TestConfigOpenInvalid(t *testing.T) {
	var err error
	testConfig, err = config.Open("/wtf/dude")
	assert.Error(t, err)
}

func TestConfigOpenValid(t *testing.T) {
	var validPath = filepath.Join(
		os.Getenv("GOPATH"), "src", "github.com", "bloGol", "bloGol", "configs",
	)

	var err error
	testConfig, err = config.Open(validPath)
	assert.NoError(t, err)
}
