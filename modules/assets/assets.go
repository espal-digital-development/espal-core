package assets

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
)

type storage interface {
	Set(key string, value []byte) error
}

// Config Assets provider object.
type Config struct {
	PublicRootFilesPath string
	ImagesPath          string
	StylesheetsPath     string
	JavaScriptPath      string
}

// Assets provider object.
type Assets struct {
	publicRootFilesPath string
	imagesPath          string
	stylesheetsPath     string
	javaScriptPath      string
}

// SetPublicRootFiles loads in all public root files data into the given storage.
func (a *Assets) SetPublicRootFiles(storage storage) error {
	if a.publicRootFilesPath == "" {
		return nil
	}
	return a.loadAllFiles(a.publicRootFilesPath, nil, storage)
}

// SetImages loads in all image data into the given storage.
func (a *Assets) SetImages(storage storage) error {
	if a.imagesPath == "" {
		return nil
	}
	return a.loadAllFiles(a.imagesPath, []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg"}, storage)
}

// SetStylesheets loads in all stylesheet data into the given storage.
func (a *Assets) SetStylesheets(storage storage) error {
	if a.stylesheetsPath == "" {
		return nil
	}
	return a.loadAllFiles(a.stylesheetsPath, []string{".css", ".map"}, storage)
}

// SetJavaScript loads in all JavaSript data into the given storage.
func (a *Assets) SetJavaScript(storage storage) error {
	if a.javaScriptPath == "" {
		return nil
	}
	return a.loadAllFiles(a.javaScriptPath, []string{".js", ".map"}, storage)
}

func (a *Assets) loadAllFiles(subjectPath string, extensions []string, storage storage) error {
	subjectPathSlash := subjectPath + "/"
	return filepath.Walk(subjectPath, func(path string, info os.FileInfo, err error) error {
		// Walk uses backward slashes on Windows, so replace them with forward
		// Need forward slashes, because assets are being served based on forward slashes
		path = strings.ReplaceAll(path, "\\", "/")
		if err != nil {
			return errors.Trace(err)
		}
		if info.IsDir() {
			return nil
		}
		if extensions != nil {
			ext := filepath.Ext(path)
			var validExt bool
			for k := range extensions {
				if extensions[k] == ext {
					validExt = true
					break
				}
			}
			if !validExt {
				return nil
			}
		}
		bytes, readErr := ioutil.ReadFile(path)
		if readErr != nil {
			return errors.Trace(readErr)
		}
		if setErr := storage.Set(strings.TrimPrefix(path, subjectPathSlash), bytes); err != nil {
			return errors.Trace(setErr)
		}
		return nil
	})
}

// New returns a new instance of Assets.
func New(config *Config) (*Assets, error) {
	a := &Assets{
		publicRootFilesPath: config.PublicRootFilesPath,
		imagesPath:          config.ImagesPath,
		stylesheetsPath:     config.StylesheetsPath,
		javaScriptPath:      config.JavaScriptPath,
	}
	return a, nil
}
