package profile_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/pprof/profile"
)

func TestNew(t *testing.T) {
	profile := profile.New()
	if profile == nil {
		t.Fatal("expected profile to not be nil")
	}
}
