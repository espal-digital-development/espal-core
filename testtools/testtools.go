package testtools

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/juju/errors"
)

// This package is solely meant to assist in testing functions.
// Don't globally call them inside normal package code-flows.

const hashLength = 8

// RequestNewTempDir will spawn a new temporary directory for tests to use.
func RequestNewTempDir(t *testing.T) string {
	tmpDir, err := filePathAbs(os.TempDir() + "/espal-tests-" + randomString(hashLength))
	if err != nil {
		t.Fatal(err)
	}
	return tmpDir
}

func filePathAbs(path string) (string, error) {
	path, err := filepath.Abs(strings.TrimRight(strings.TrimRight(path, "/"), "\\"))
	if err != nil {
		return "", errors.Trace(err)
	}
	return strings.ReplaceAll(path, "\\", "/"), nil
}

// EqString checks if the given strings are equal.
func EqString(t *testing.T, subject string, have string, want string) {
	if want != have {
		t.Fatalf("Expected `%s` to be `%s` but got `%s`", subject, want, have)
	}
}

// EqBool checks if the given booleans are equal.
func EqBool(t *testing.T, subject string, have bool, want bool) {
	if want != have {
		t.Fatalf("Expected `%s` to be `%v` but got `%v`", subject, want, have)
	}
}

// EqInt checks if the given integers are equal.
func EqInt(t *testing.T, subject string, have int, want int) {
	if want != have {
		t.Fatalf("Expected `%s` to be `%d` but got `%d`", subject, want, have)
	}
}

// EqDuration checks if the given durations are equal.
func EqDuration(t *testing.T, subject string, have time.Duration, want time.Duration) {
	if want != have {
		t.Fatalf("Expected `%s` to be `%v` but got `%v`", subject, want, have)
	}
}

// EqError checks if the given errors are equal.
func EqError(t *testing.T, subject string, have error, want error) {
	if want.Error() != have.Error() {
		t.Fatalf("Expected `%s` to be `%v` but got `%v`", subject, want, have)
	}
}
