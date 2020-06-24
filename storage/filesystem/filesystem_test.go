package filesystem_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/storage/filesystem"
	"github.com/espal-digital-development/espal-core/testtools"
)

const testKey = "testdata"

func newStorage(t *testing.T) storage.Storage {
	tmpDir := testtools.RequestNewTempDir(t)
	if err := os.MkdirAll(tmpDir, 0700); err != nil {
		t.Fatal(err)
	}
	storage, err := filesystem.New(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	return storage
}

func TestNonExistentPath(t *testing.T) {
	storage, err := filesystem.New("nothing--------")
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if storage != nil {
		t.Fatal("Storage should be nil when an error was thrown")
	}
}

func TestExists(t *testing.T) {
	storage := newStorage(t)
	ok := storage.Exists(testKey)
	if ok {
		t.Fatal("Entry should not exist")
	}
	expectedData := []byte("testdata")
	err := storage.Set(testKey, expectedData)
	if err != nil {
		t.Fatal(err)
	}
	ok = storage.Exists(testKey)
	if !ok {
		t.Fatal("Entry should exist")
	}
}

func TestSetGet(t *testing.T) {
	storage := newStorage(t)
	expectedData := []byte("testdata")
	err := storage.Set(testKey, expectedData)
	if err != nil {
		t.Fatal(err)
	}
	data, ok, err := storage.Get(testKey)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Expected data to be found")
	}
	if !bytes.Equal(data, expectedData) {
		t.Fatalf("Expected data to be `%v` but got `%v`", expectedData, data)
	}
}

func TestGetNonExistent(t *testing.T) {
	storage := newStorage(t)
	data, ok, err := storage.Get("testdata")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("Expected no data to be found")
	}
	if len(data) > 0 {
		t.Fatal("Data should be empty for non-existent entries")
	}
}

func TestDelete(t *testing.T) {
	storage := newStorage(t)
	key := "testdata"
	expectedData := []byte("testdata")
	err := storage.Set(key, expectedData)
	if err != nil {
		t.Fatal(err)
	}
	data, ok, err := storage.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Expected data to be found")
	}
	if !bytes.Equal(data, expectedData) {
		t.Fatalf("Expected data to be `%v` but got `%v`", expectedData, data)
	}
	err = storage.Delete(key)
	if err != nil {
		t.Fatal(err)
	}
	data, ok, err = storage.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("Expected no data to be found")
	}
	if len(data) > 0 {
		t.Fatal("Data should be empty for non-existent entries")
	}
}
