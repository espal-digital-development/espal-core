package posts_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/routes/pprof/posts"
)

func TestNew(t *testing.T) {
	posts := posts.New("pprof_prefix")
	if posts == nil {
		t.Fatal("expected posts to not be nil")
	}
}
