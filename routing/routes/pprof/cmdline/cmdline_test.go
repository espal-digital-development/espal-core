package cmdline_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/pprof/cmdline"
)

func TestNew(t *testing.T) {
	cmdLine := cmdline.New()
	if cmdLine == nil {
		t.Fatal("expected cmdLine to not be nil")
	}
}
