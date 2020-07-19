package repositories_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/modules/repositories"
)

func TestNew(t *testing.T) {
	repositories, err := repositories.New()
	if err != nil {
		t.Fatal(err)
	}
	if repositories == nil {
		t.Fatal("expected repositories to not be nil")
	}
}
