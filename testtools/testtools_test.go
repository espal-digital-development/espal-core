package testtools_test

import (
	"errors"
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/testtools"
)

func TestRequestNewTempDir(t *testing.T) {
	tmpDir := testtools.RequestNewTempDir(t)
	if tmpDir == "" {
		t.Fatal("tmpDir should not be empty")
	}
}

func TestEqBool(t *testing.T) {
	have := false
	testtools.EqBool(t, "subject", have, have)
}

func TestEqDuration(t *testing.T) {
	have := time.Second
	testtools.EqDuration(t, "subject", have, have)
}

func TestEqError(t *testing.T) {
	have := errors.New("test")
	testtools.EqError(t, "subject", have, have)
}

func TestEqInt(t *testing.T) {
	have := 42
	testtools.EqInt(t, "subject", have, have)
}

func TestEqString(t *testing.T) {
	have := "a"
	testtools.EqString(t, "subject", have, have)
}
