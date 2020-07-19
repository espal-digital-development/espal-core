package health_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/health"
)

func TestNew(t *testing.T) {
	health := health.New()
	if health == nil {
		t.Fatal("expected health to not be nil")
	}
}
