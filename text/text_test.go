package text_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/text"
)

func TestRandomString(t *testing.T) {
	for i := 1; i <= 2^12; i++ {
		str := text.RandomString(i)
		if len(str) != i {
			t.Errorf("expected to %d long, but got %d characters on `%s`", i, i, str)
		}
	}
}

func TestLowerFirst(t *testing.T) {
	word := text.LowerFirst("Test")
	if word != "test" {
		t.Fatalf("word should be `test` but got `%s`", word)
	}
}

func TestLowerFirstEmpty(t *testing.T) {
	text.LowerFirst("")
}
