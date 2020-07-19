package config_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/modules/config"
)

func TestNew(t *testing.T) {
	config, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	if config == nil {
		t.Fatal("expected config to not be nil")
	}
}
