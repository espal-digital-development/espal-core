package database_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/database"
	_ "github.com/lib/pq"
)

// TODO :: 77 This makes no sense; sql doesn't really fail on connect, only
// later when you actually try to do queries. This should actually just be
// an ignored test as it tests nothing at all.

const (
	driverName = "postgres"
	dsn        = "---"
)

func TestNew(t *testing.T) {
	database := database.New()
	if database == nil {
		t.Fatal("database should not be nil")
	}
}

func TestOpen(t *testing.T) {
	database := database.New()
	if err := database.Open(driverName, dsn); err != nil {
		t.Fatal(err)
	}
}

func TestClose(t *testing.T) {
	database := database.New()
	if err := database.Open(driverName, dsn); err != nil {
		t.Fatal(err)
	}
	if err := database.Close(); err != nil {
		t.Fatal(err)
	}
}
