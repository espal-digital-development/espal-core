package routes_test

import (
	"errors"
	"testing"

	"github.com/espal-digital-development/espal-core/modules/routes"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

var errIterate = errors.New("mock iteration error")

type mockHandler struct{}

func (h *mockHandler) Handle(contexts.Context) {}

func TestRegister(t *testing.T) {
	routes, err := routes.New(&routes.Config{
		Entries: map[string]routes.Handler{
			"/test": &mockHandler{},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if routes == nil {
		t.Fatal("expected routes to not be nil")
	}
}

func TestIterate(t *testing.T) {
	r, err := routes.New(&routes.Config{
		Entries: map[string]routes.Handler{
			"/test":  &mockHandler{},
			"/test2": &mockHandler{},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	var c int
	err = r.Iterate(func(path string, h routes.Handler) error {
		c++
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if c != 2 {
		t.Fatalf("iterate should've return the 2 registered routes, but returned %d", c)
	}
}

func TestIterateEmpty(t *testing.T) {
	r, err := routes.New(&routes.Config{})
	if err != nil {
		t.Fatal(err)
	}
	var c int
	err = r.Iterate(func(path string, h routes.Handler) error {
		c++
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if c != 0 {
		t.Fatalf("iterate should've return the 0 registered routes, but returned %d", c)
	}
}

func TestIterateError(t *testing.T) {
	r, err := routes.New(&routes.Config{
		Entries: map[string]routes.Handler{
			"/test": &mockHandler{},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = r.Iterate(func(path string, h routes.Handler) error {
		return errIterate
	})
	if err == nil {
		t.Fatal("iterate should return an error")
	}
}
