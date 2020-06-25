package currencies_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/languages/languagesmock"
	"github.com/espal-digital-development/espal-core/testtools"
	"github.com/juju/errors"
)

const (
	euroMemberCountries = "Euro Member Countries"
)

var (
	availableLanguages = map[string]*languagesmock.DataMock{
		"en": {
			IDFunc: func() uint16 {
				return 1
			},
		},
		"nl": {
			IDFunc: func() uint16 {
				return 2
			},
		},
	}
	languagesRepository *languagesmock.RepositoryMock
)

func initMocks() {
	languagesRepository = &languagesmock.RepositoryMock{
		ByCodeFunc: func(code string) (languages.Data, error) {
			if _, ok := availableLanguages[code]; !ok {
				return availableLanguages["en"], nil
			}
			return availableLanguages[code], nil
		},
	}
}

func TestAll(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	all := currencies.All()
	if len(all) != 186 {
		t.Fatal("All returns an incorrect amount")
	}

	twoCharUpper, err := regexp.Compile(`^[A-Z]{3}$`)
	if err != nil {
		t.Fatal(err)
	}

	for id := range all {
		c, err := currencies.ByID(id)
		if err != nil {
			t.Fatal(err)
		}
		code := c.Code()
		if !twoCharUpper.MatchString(code) {
			t.Errorf("`%s` is an invalid code", code)
		}
		name := c.EnglishName()
		if len(name) < 4 {
			t.Errorf("name of `%s` shorter than 4 characters", name)
		}
	}
}

func TestByID(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	_, err = currencies.ByID(33)
	if err != nil {
		t.Fatal(err)
	}

	_, err = currencies.ByID(0)
	if err == nil {
		t.Fatal("Calling unexisting should give an error")
	}
}

func TestByCode(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	_, err = currencies.ByCode("EUR")
	if err != nil {
		t.Fatal(err)
	}

	_, err = currencies.ByCode("00")
	if err == nil {
		t.Fatal("Calling unexisting should give an error")
	}
}

func TestCurrencySymbol(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	currency, err := currencies.ByCode("EUR")
	if err != nil {
		t.Fatal(err)
	}

	if currency.Symbol() != "€" {
		t.Fatalf("args")
	}
}

func TestCurrencyID(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	currency, err := currencies.ByCode("EUR")
	if err != nil {
		t.Fatal(err)
	}
	if currency.ID() == 0 {
		t.Fatal("ID should not be 0")
	}
}

func TestCurrencyHasTranslation(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	es, err := languagesRepository.ByCode("es")
	if err != nil {
		t.Fatal(err)
	}
	currency, err := currencies.ByCode("EUR")
	if err != nil {
		t.Fatal(err)
	}
	trans, ok := currency.HasTranslation(es.ID())
	if !ok {
		t.Fatal("Failed to find the Spanish translations for Euro")
	}
	if strings.TrimSpace(trans) == "" {
		t.Fatal("Translation returned empty")
	}
}

func TestCurrencyTranslate(t *testing.T) {
	initMocks()
	currencies, err := currencies.New(languagesRepository)
	if err != nil {
		t.Fatal(err)
	}

	es, err := languagesRepository.ByCode("es")
	if err != nil {
		t.Fatal(err)
	}
	en, err := languagesRepository.ByCode("en")
	if err != nil {
		t.Fatal(err)
	}
	c, err := currencies.ByCode("EUR")
	if err != nil {
		t.Fatal(err)
	}
	if c.Translate(es.ID()) != euroMemberCountries {
		t.Fatalf("Expected `%s` but got `%s`", euroMemberCountries, c.Translate(es.ID()))
	}
	if c.Translate(en.ID()) != euroMemberCountries {
		t.Fatalf("Expected `%s` but got `%s`", euroMemberCountries, c.Translate(en.ID()))
	}
	if c.Translate(0) != euroMemberCountries {
		t.Fatalf("Expected `%s` but got `%s`", euroMemberCountries, c.Translate(0))
	}
}

func TestEnglishLanguageFailure(t *testing.T) {
	initMocks()
	languageError := errors.New("languageError")
	languagesRepository.ByCodeFunc = func(code string) (languages.Data, error) {
		return nil, languageError
	}
	currencies, err := currencies.New(languagesRepository)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "languageError", err, languageError)
	if currencies != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestEnglishLanguageFailureInLoadTranslations(t *testing.T) {
	initMocks()
	var callCounter int
	languageError := errors.New("languageError")
	languagesRepository.ByCodeFunc = func(code string) (languages.Data, error) {
		callCounter++
		if callCounter == 1 {
			if _, ok := availableLanguages[code]; !ok {
				return availableLanguages["en"], nil
			}
			return availableLanguages[code], nil
		}
		return nil, languageError
	}
	currencies, err := currencies.New(languagesRepository)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "languageError", err, languageError)
	if currencies != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}
