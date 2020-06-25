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

// All gets all embedded l.
func (l *Languages) All() map[uint16]Data {
	return l.all
}

// ByID looks for and returns the associated Language for the given id.
func (l *Languages) ByID(id uint16) (Data, error) {
	if _, ok := l.byID[id]; !ok {
		return &Language{}, errors.Errorf("no language found with given id `%d`", id)
	}
	return l.byID[id], nil
}

// ByCode looks for and returns the associated Language for the given code.
func (l *Languages) ByCode(code string) (Data, error) {
	if _, ok := l.entries[code]; !ok {
		return &Language{}, errors.Errorf("no language found with given code `%s`", code)
	}
	return l.entries[code], nil
}

func (l *Languages) loadTranslations() error {
	files, err := languagesdata.AssetDir("_data")
	if err != nil {
		return errors.Trace(err)
	}
	for k := range files {
		if len(files[k]) == 0 || files[k][0] == '.' {
			continue
		}
		language := l.entries[files[k]]
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
				targetLanguage, err := l.ByID(count)
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
func (l *Languages) EnglishLocaleID() uint16 {
	return data["en"].id
}

// New returns new a *Languages repository instance.
func New() (*Languages, error) {
	l := &Languages{
		byID:    make(map[uint16]*Language, len(data)),
		all:     make(map[uint16]Data, len(data)),
		entries: make(map[string]*Language, len(data)),
	}

	english := data["en"]
	for code := range data {
		language := data[code]
		language.englishLocaleID = english.ID()
		language.translations = map[uint16]string{}

		l.byID[language.ID()] = &language
		l.all[language.ID()] = &language
		l.entries[code] = &language
	}

	if err := l.loadTranslations(); err != nil {
		return nil, errors.Trace(err)
	}

	return l, nil
}
