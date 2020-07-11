package modules_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/modules"
)

func TestNew(t *testing.T) {
	config := &modules.Config{}
	module, err := modules.New(config)
	if err != nil {
		t.Fatal(err)
	}
	if module == nil {
		t.Fatal("expected module to not be nil")
	}
}
