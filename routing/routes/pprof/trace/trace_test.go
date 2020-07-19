package trace_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/pprof/trace"
)

func TestNew(t *testing.T) {
	trace := trace.New()
	if trace == nil {
		t.Fatal("expected trace to not be nil")
	}
}
