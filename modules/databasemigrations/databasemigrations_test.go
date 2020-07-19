package databasemigrations_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/modules/databasemigrations"
)

func TestNew(t *testing.T) {
	databaseMigrations, err := databasemigrations.New()
	if err != nil {
		t.Fatal(err)
	}
	if databaseMigrations == nil {
		t.Fatal("expected databaseMigrations to not be nil")
	}
}
