package assets_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/modules/assets"
)

type storageMock struct {
	failSet bool
}

func (m *storageMock) Set(key string, value []byte) error {
	if m.failSet {
		return errors.New("fake error")
	}
	return nil
}

func TestNew(t *testing.T) {
	assets, err := assets.New(&assets.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if assets == nil {
		t.Fatal("expected assets to not be nil")
	}
}

func TestEmptySetters(t *testing.T) {
	assets, err := assets.New(&assets.Config{})
	if err != nil {
		t.Fatal(err)
	}
	storageMock := &storageMock{}
	if err := assets.SetPublicRootFiles(storageMock); err != nil {
		t.Fatal(err)
	}
	if err := assets.SetImages(storageMock); err != nil {
		t.Fatal(err)
	}
	if err := assets.SetStylesheets(storageMock); err != nil {
		t.Fatal(err)
	}
	if err := assets.SetJavaScript(storageMock); err != nil {
		t.Fatal(err)
	}
}

func TestLoaders(t *testing.T) {
	random := rand.NewSource(time.Now().UnixNano())
	tmpDir := os.TempDir()
	hashDirPrefix := fmt.Sprintf("%s%d_", tmpDir, random.Int63())
	config := &assets.Config{
		PublicRootFilesPath: filepath.FromSlash(hashDirPrefix + "root"),
		ImagesPath:          filepath.FromSlash(hashDirPrefix + "images"),
		StylesheetsPath:     filepath.FromSlash(hashDirPrefix + "css"),
		JavaScriptPath:      filepath.FromSlash(hashDirPrefix + "js"),
	}
	assets, err := assets.New(config)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(config.PublicRootFilesPath, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(config.ImagesPath, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(config.StylesheetsPath, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(config.JavaScriptPath, 0700); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.FromSlash(config.StylesheetsPath+"/test.css"), []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.FromSlash(config.StylesheetsPath+"/no_css.txt"), []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	storageMock := &storageMock{}
	if err := assets.SetPublicRootFiles(storageMock); err != nil {
		t.Fatal(err)
	}
	if err := assets.SetImages(storageMock); err != nil {
		t.Fatal(err)
	}
	if err := assets.SetStylesheets(storageMock); err != nil {
		t.Fatal(err)
	}
	if err := assets.SetJavaScript(storageMock); err != nil {
		t.Fatal(err)
	}
}

func TestStorageSetterFailure(t *testing.T) {
	tmpDir := os.TempDir()
	assets, err := assets.New(&assets.Config{
		StylesheetsPath: tmpDir,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.FromSlash(tmpDir+"test.css"), []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	storageMock := &storageMock{
		failSet: true,
	}
	// TODO :: Coverage doesn't seem to catch the walker's return, but it does catch it on the next cycle.
	if err := assets.SetStylesheets(storageMock); err == nil {
		t.Fatal("the setter should fail")
	}
}
