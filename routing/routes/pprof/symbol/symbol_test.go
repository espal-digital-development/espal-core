package symbol_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/pprof/symbol"
)

func TestNew(t *testing.T) {
	symbol := symbol.New()
	if symbol == nil {
		t.Fatal("expected symbol to not be nil")
	}
}
