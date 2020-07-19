package overview_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/pprof/overview"
)

func TestNew(t *testing.T) {
	overview := overview.New()
	if overview == nil {
		t.Fatal("expected overview to not be nil")
	}
}
