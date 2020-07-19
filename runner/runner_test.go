package runner_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/runner"
)

func TestNew(t *testing.T) {
	runner, err := runner.New()
	if err != nil {
		t.Fatal(err)
	}
	if runner == nil {
		t.Fatal("expected runner to not be nil")
	}
}
