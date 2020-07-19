package translations_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/modules/translations"
)

func TestNew(t *testing.T) {
	translations, err := translations.New(&translations.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if translations == nil {
		t.Fatal("expected translations to not be nil")
	}
}
