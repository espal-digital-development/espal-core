package languages

import (
	"bytes"
	"io"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/languages/languagesdata"
	"github.com/juju/errors"
)

var _ Repository = &Languages{}

// Repository contains a full Language repository.
type Repository interface {
	All() map[uint16]Data
	ByID(id uint16) (Data, error)
	ByCode(code string) (Data, error)
	EnglishLocaleID() uint16
}

// Languages contains a full Language repository.
type Languages struct {
	entries map[string]*Language
	byID    map[uint16]*Language
	all     map[uint16]Data
}

// All gets all embedded languages.
func (languages *Languages) All() map[uint16]Data {
	return languages.all
}

// ByID looks for and returns the associated Language for the given id.
func (languages *Languages) ByID(id uint16) (Data, error) {
	if _, ok := languages.byID[id]; !ok {
		return &Language{}, errors.Errorf("no language found with given id `%d`", id)
	}
	return languages.byID[id], nil
}

// ByCode looks for and returns the associated Language for the given code.
func (languages *Languages) ByCode(code string) (Data, error) {
	if _, ok := languages.entries[code]; !ok {
		return &Language{}, errors.Errorf("no language found with given code `%s`", code)
	}
	return languages.entries[code], nil
}

func (languages *Languages) loadTranslations() error {
	files, err := languagesdata.AssetDir("_data")
	if err != nil {
		return errors.Trace(err)
	}
	for k := range files {
		if len(files[k]) == 0 || files[k][0] == '.' {
			continue
		}
		language := languages.entries[files[k]]
		data, err := languagesdata.Asset("_data/" + files[k])
		if err != nil {
			return errors.Trace(err)
		}

		buf := bytes.NewBuffer(data)
		var count uint16 = 1
		for {
			t, err := buf.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				return errors.Trace(err)
			}
			if len(t) > 0 && t[0] != '\n' {
				targetLanguage, err := languages.ByID(count)
				if err != nil {
					return errors.Trace(err)
				}
				language.SetTranslation(targetLanguage.ID(), strings.Trim(string(t), "\n"))
			}
			count++
		}
	}
	return nil
}

// EnglishLocaleID returns the actual default English locale ID number.
func (languages *Languages) EnglishLocaleID() uint16 {
	return data["en"].id
}

// New returns new a *Languages repository instance.
func New() (*Languages, error) {
	english := data["en"]
	languages := &Languages{
		byID:    make(map[uint16]*Language, len(data)),
		all:     make(map[uint16]Data, len(data)),
		entries: make(map[string]*Language, len(data)),
	}

	for code := range data {
		language := data[code]
		language.englishLocaleID = english.ID()
		language.translations = map[uint16]string{}

		languages.byID[language.ID()] = &language
		languages.all[language.ID()] = &language
		languages.entries[code] = &language
	}

	if err := languages.loadTranslations(); err != nil {
		return nil, errors.Trace(err)
	}

	return languages, nil
}
