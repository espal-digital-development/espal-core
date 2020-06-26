package currencies

import (
	"bytes"
	"io"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/currencies/currenciesdata"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/juju/errors"
)

var _ Repository = &Currencies{}

// Repository represents a Currencies repository.
type Repository interface {
	All() map[uint]Data
	ByID(id uint) (Data, error)
	ByCode(code string) (Data, error)
}

// Currencies contains a full Currency repository.
// CLDR bcp47 currencies ISO4217.
type Currencies struct {
	languagesRepository languages.Repository
	entries             map[string]*Currency
	byID                map[uint]*Currency
	all                 map[uint]Data
}

// All gets all embedded c.
func (c *Currencies) All() map[uint]Data {
	return c.all
}

// ByID looks for and returns the associated Currency for the given id.
func (c *Currencies) ByID(id uint) (Data, error) {
	if _, ok := c.byID[id]; !ok {
		return &Currency{}, errors.Errorf("no currency found with given id `%d`", id)
	}
	return c.byID[id], nil
}

// ByCode looks for and returns the associated Currency for the given code.
func (c *Currencies) ByCode(code string) (Data, error) {
	if _, ok := c.entries[code]; !ok {
		return &Currency{}, errors.Errorf("no currency found with given code `%s`", code)
	}
	return c.entries[code], nil
}

func (c *Currencies) loadTranslations() error {
	files, err := currenciesdata.AssetDir("_data")
	if err != nil {
		return errors.Trace(err)
	}
	for k := range files {
		if len(files[k]) == 0 || files[k][0] == '.' {
			continue
		}
		data, err := currenciesdata.Asset("_data/" + files[k])
		if err != nil {
			return errors.Trace(err)
		}

		language, err := c.languagesRepository.ByCode(files[k])
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
				return errors.Errorf("currency read failure in file `%s` for line %d", files[k], count)
			}
			checkCurrency, err := c.ByCode(parts[0])
			if err != nil {
				return errors.Errorf("unknown Currency code `%s` in file `%s` for line %d", parts[0], files[k], count)
			}
			currency := c.entries[checkCurrency.Code()]
			if parts[1] == currency.EnglishName() {
				return errors.Errorf("duplicate Currency English translation for code `%s` in file `%s` for line %d",
					parts[0], files[k], count)
			}
			currency.SetTranslation(language.ID(), parts[1])

			count++
		}
	}
	return nil
}

// New returns new a *Currencies repository instance.
func New(languagesRepository languages.Repository) (*Currencies, error) {
	c := &Currencies{
		languagesRepository: languagesRepository,
		byID:                make(map[uint]*Currency, len(data)),
		all:                 make(map[uint]Data, len(data)),
		entries:             make(map[string]*Currency, len(data)),
	}

	english, err := languagesRepository.ByCode("en")
	if err != nil {
		return nil, errors.Trace(err)
	}

	// Fill the currenciesByID shortcut map
	for k := range data {
		currency := data[k]
		currency.translations = make(map[uint16]string)
		currency.englishLocaleID = english.ID()
		c.byID[currency.id] = &currency
		c.all[currency.id] = &currency
		c.entries[currency.Code()] = &currency
	}

	if err := c.loadTranslations(); err != nil {
		return nil, errors.Trace(err)
	}

	return c, nil
}
