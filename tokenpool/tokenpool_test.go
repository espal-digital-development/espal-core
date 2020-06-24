package tokenpool_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/tokenpool"
)

const (
	defaultExpiration      = time.Minute * 8
	defaultCleanupInterval = time.Second * 10
)

func TestNew(t *testing.T) {
	pool := tokenpool.New(defaultExpiration, defaultCleanupInterval)
	if pool == nil {
		t.Fatal("Expected pool to not be nil")
	}
}

func TestRequestToken(t *testing.T) {
	pool := tokenpool.New(defaultExpiration, defaultCleanupInterval)
	token, err := pool.RequestToken()
	if err != nil {
		t.Fatal(err)
	}
	if token <= 0 {
		t.Fatal("Expected token to exist")
	}
}

func TestValidate(t *testing.T) {
	pool := tokenpool.New(defaultExpiration, defaultCleanupInterval)
	token, err := pool.RequestToken()
	if err != nil {
		t.Fatal(err)
	}
	ok := pool.Validate(token)
	if !ok {
		t.Fatal("Expected token to exist")
	}
	ok = pool.Validate(token)
	if ok {
		t.Fatal("Expected token to been consumed")
	}
}

func TestCount(t *testing.T) {
	pool := tokenpool.New(defaultExpiration, defaultCleanupInterval)
	for i := 1; i <= 5; i++ {
		_, err := pool.RequestToken()
		if err != nil {
			t.Fatal(err)
		}
		if uint(i) != pool.Count() {
			t.Fatalf("Expected Count to be %d", i)
		}
	}
}

func TestCleanup(t *testing.T) {
	pool := tokenpool.New(time.Nanosecond, time.Millisecond)
	_, err := pool.RequestToken()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 2)
	if pool.Count() != 0 {
		t.Fatal("Count should return 0 after Cleanup")
	}
}
