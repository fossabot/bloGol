package db

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	if err := Open("/wtf/dude"); err == nil {
		t.Error("Except error, but got nil")
	}

	var validPath = filepath.Join(
		os.Getenv("GOPATH"), "src", "github.com", "bloGol", "bloGol", "test", "development.db",
	)

	if err := Open(validPath); err != nil {
		t.Error(err.Error())
	}
}
