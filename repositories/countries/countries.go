package countries

import (
	"bytes"
	"io"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/countries/countriesdata"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/juju/errors"
)

var _ Repository = &Countries{}

// Repository represents a Country repository.
type Repository interface {
	All() map[uint16]Data
	ByID(id uint16) (Data, error)
	ByCode(code string) (Data, error)
}

// Countries contains a full Country repository.
type Countries struct {
	languagesRepository languages.Repository
	entries             map[string]*Country
	byID                map[uint16]*Country
	all                 map[uint16]Data
}

func (countriesRepository *Countries) loadTranslations() error {
	files, err := countriesdata.AssetDir("_data")
	if err != nil {
		return errors.Trace(err)
	}
	for k := range files {
		if len(files[k]) == 0 || files[k][0] == '.' {
			continue
		}
		data, err := countriesdata.Asset("_data/" + files[k])
		if err != nil {
			return errors.Trace(err)
		}

		language, err := countriesRepository.languagesRepository.ByCode(files[k])
		if err != nil {
			return errors.Trace(err)
		}

		buf := bytes.NewBuffer(data)
		var count uint16 = 1
		for {
			t, err := buf.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				return errors.Trace(err)
			}

			parts := strings.Split(strings.Trim(t, "\n"), "\t")
			if len(parts) != 2 {
				return errors.Errorf("country read failure in file `%s` for line %d", files[k], count)
			}
			checkCountry, err := countriesRepository.ByCode(parts[0])
			if err != nil {
				return errors.Errorf("unknown Country code `%s` in file `%s` for line %d", parts[0], files[k], count)
			}
			country := countriesRepository.entries[checkCountry.Code()]
			if parts[1] == country.EnglishName() {
				return errors.Errorf("duplicate Country English translation for code `%s` in file `%s` for line %d", parts[0], files[k], count)
			}
			country.SetTranslation(language.ID(), parts[1])

			count++
		}
	}
	return nil
}

func (countriesRepository *Countries) setAliases() error {
	aliases := map[string]string{
		"ar_001":  "ar",
		"de_AT":   "de",
		"de_CH":   "de",
		"es_419":  "es",
		"es_ES":   "es",
		"es_MX":   "es",
		"fa_AF":   "fa",
		"fr_CA":   "fr",
		"fr_CH":   "fr",
		"nl_BE":   "nl",
		"pt_BR":   "pt",
		"pt_PT":   "pt",
		"ro_MD":   "ro",
		"sw_CD":   "sw",
		"zh_Hans": "zh",
		"zh_Hant": "zh",
	}
	for source, target := range aliases {
		sourceLanguage, err := countriesRepository.languagesRepository.ByCode(source)
		if err != nil {
			return errors.Trace(err)
		}
		targetLanguage, err := countriesRepository.languagesRepository.ByCode(target)
		if err != nil {
			return errors.Trace(err)
		}
		for k := range countriesRepository.entries {
			targetTranslation := countriesRepository.entries[k].Translate(targetLanguage.ID())
			countriesRepository.entries[k].SetTranslation(sourceLanguage.ID(), targetTranslation)
		}
	}
	return nil
}

// All gets all embedded Country entries.
func (countriesRepository *Countries) All() map[uint16]Data {
	return countriesRepository.all
}

// ByID looks for and returns the associated Country for the given id.
func (countriesRepository *Countries) ByID(id uint16) (Data, error) {
	if _, ok := countriesRepository.byID[id]; !ok {
		return &Country{}, errors.Errorf("no country found with given id `%d`", id)
	}
	return countriesRepository.byID[id], nil
}

// ByCode looks for and returns the associated Country for the given code.
func (countriesRepository *Countries) ByCode(code string) (Data, error) {
	if _, ok := countriesRepository.entries[code]; !ok {
		return &Country{}, errors.Errorf("no country found with given code `%s`", code)
	}
	return countriesRepository.entries[code], nil
}

// New returns new a *Countries instance.
func New(languagesRepository languages.Repository, loadTranslations bool) (*Countries, error) {
	countries := &Countries{
		languagesRepository: languagesRepository,
		byID:                make(map[uint16]*Country, len(data)),
		all:                 make(map[uint16]Data, len(data)),
		entries:             make(map[string]*Country, len(data)),
	}

	english, err := languagesRepository.ByCode("en")
	if err != nil {
		return nil, errors.Trace(err)
	}

	for code := range data {
		country := data[code]
		country.englishLocaleID = english.ID()
		country.translations = map[uint16]string{}

		countries.byID[country.ID()] = &country
		countries.all[country.ID()] = &country
		countries.entries[code] = &country
	}

	if loadTranslations {
		if err := countries.loadTranslations(); err != nil {
			return nil, errors.Trace(err)
		}
		if err := countries.setAliases(); err != nil {
			return nil, errors.Trace(err)
		}
	}

	return countries, nil
}
