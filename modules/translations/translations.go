package translations

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
)

// Config Translations provider object.
type Config struct {
	Path string
}

// Translations provider object.
type Translations struct {
	path string
}

// GetAll fetches the translations file path per language.
func (t *Translations) GetAll() (map[string]string, error) {
	if t.path == "" {
		return nil, nil
	}

	translations := map[string]string{}
	files, err := ioutil.ReadDir(t.path)
	if err != nil {
		return nil, errors.Trace(err)
	}
	for k := range files {
		if files[k].IsDir() {
			continue
		}
		ext := filepath.Ext(files[k].Name())
		if ext != ".yml" {
			continue
		}
		language := strings.TrimSuffix(files[k].Name(), ext)
		translations[language] = filepath.FromSlash(t.path + "/" + files[k].Name())
	}

	return translations, nil
}

// New returns a new instance of Translations.
func New(config *Config) (*Translations, error) {
	t := &Translations{
		path: config.Path,
	}
	return t, nil
}
