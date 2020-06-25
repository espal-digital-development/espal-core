package countries_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/languages/languagesmock"
	"github.com/juju/errors"
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
	countries, err := countries.New(languagesRepository, true)
	if err != nil {
		t.Fatal(err)
	}

	all := countries.All()
	if len(all) != 248 {
		t.Fatal("All returns an incorrect amount")
	}

	twoCharUpper, err := regexp.Compile(`^[A-Z]{2}$`)
	if err != nil {
		t.Fatal(err)
	}

	for id := range all {
		c, err := countries.ByID(id)
		if err != nil {
			t.Fatal(err)
		}
		code := c.Code()
		if !twoCharUpper.MatchString(code) {
			t.Errorf("`%s` is an invalid code", code)
		}
		name := c.EnglishName()
		if len(name) < 3 {
			t.Errorf("name of `%s` shorter than 3 characters", name)
		}
	}
}

func TestByID(t *testing.T) {
	initMocks()
	countries, err := countries.New(languagesRepository, true)
	if err != nil {
		t.Fatal(err)
	}

	c, err := countries.ByID(528)
	if err != nil {
		t.Fatal(err)
	}
	if c.ID() != 528 {
		t.Errorf("expected ID `528` but got `%d`", c.ID())
	}

	_, err = countries.ByID(0)
	if err == nil {
		t.Error("calling unexisting should give an error")
	}
}

func TestByCode(t *testing.T) {
	initMocks()
	countries, err := countries.New(languagesRepository, true)
	if err != nil {
		t.Fatal(err)
	}

	c, err := countries.ByCode("NL")
	if err != nil {
		t.Fatal(err)
	}
	if c.Code() != "NL" {
		t.Errorf("expected code `NL` but got `%s`", c.Code())
	}

	_, err = countries.ByCode("00")
	if err == nil {
		t.Error("calling unexisting should give an error")
	}
}

func TestCountryID(t *testing.T) {
	initMocks()
	countries, err := countries.New(languagesRepository, true)
	if err != nil {
		t.Fatal(err)
	}

	country, err := countries.ByCode("NL")
	if err != nil {
		t.Fatal(err)
	}
	if country.ID() == 0 {
		t.Fatal("ID should not be 0")
	}
}

func TestCountryHasTranslation(t *testing.T) {
	initMocks()
	countries, err := countries.New(languagesRepository, true)
	if err != nil {
		t.Fatal(err)
	}

	es, err := languagesRepository.ByCode("es")
	if err != nil {
		t.Fatal(err)
	}
	country, err := countries.ByCode("GB")
	if err != nil {
		t.Fatal(err)
	}
	trans, ok := country.HasTranslation(es.ID())
	if !ok {
		t.Fatal("Failed to find the Spanish translations for United Kingdom")
	}
	if strings.TrimSpace(trans) == "" {
		t.Fatal("Translation returned empty")
	}
}

func TestCountryTranslate(t *testing.T) {
	initMocks()
	countries, err := countries.New(languagesRepository, true)
	if err != nil {
		t.Fatal(err)
	}

	nl, err := languagesRepository.ByCode("nl")
	if err != nil {
		t.Fatal(err)
	}
	en, err := languagesRepository.ByCode("en")
	if err != nil {
		t.Fatal(err)
	}
	c, err := countries.ByCode("GB")
	if err != nil {
		t.Fatal(err)
	}
	if c.Translate(nl.ID()) != "Verenigd Koninkrijk" {
		t.Fatalf("Expected `%s` but got `%s`", "Verenigd Koninkrijk", c.Translate(nl.ID()))
	}
	if c.Translate(en.ID()) != "United Kingdom" {
		t.Fatalf("Expected `%s` but got `%s`", "United Kingdom", c.Translate(en.ID()))
	}
	if c.Translate(0) != "United Kingdom" {
		t.Fatalf("Expected `%s` but got `%s`", "United Kingdom", c.Translate(0))
	}
}

func TestEnglishError(t *testing.T) {
	initMocks()
	languageError := errors.New("language not found")
	languagesRepository.ByCodeFunc = func(code string) (languages.Data, error) {
		return nil, languageError
	}
	_, err := countries.New(languagesRepository, true)
	if err == nil {
		t.Fatal("error should not be nil")
	}
	if languageError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", languageError.Error(), err.Error())
	}
}

func TestLoadAllLanguageError(t *testing.T) {
	initMocks()
	var errCounter uint
	languageError := errors.New("language not found")
	languagesRepository.ByCodeFunc = func(code string) (languages.Data, error) {
		if errCounter == 1 {
			return nil, languageError
		}
		errCounter++
		return availableLanguages["en"], nil
	}
	_, err := countries.New(languagesRepository, true)
	if err == nil {
		t.Fatal("error should not be nil")
	}
	if languageError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", languageError.Error(), err.Error())
	}
}

func TestSourceLanguageError(t *testing.T) {
	initMocks()
	sourceError := errors.New("source language not found")
	languagesRepository.ByCodeFunc = func(code string) (languages.Data, error) {
		if code == "ar_001" {
			return nil, sourceError
		}
		return availableLanguages["en"], nil
	}
	_, err := countries.New(languagesRepository, true)
	if err == nil {
		t.Fatal("error should not be nil")
	}
	if sourceError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", sourceError.Error(), err.Error())
	}
}

func TestTargetLanguageError(t *testing.T) {
	initMocks()
	var errCounter uint
	targetError := errors.New("target language not found")
	languagesRepository.ByCodeFunc = func(code string) (languages.Data, error) {
		if code == "zh" {
			if errCounter == 1 {
				return nil, targetError
			}
			errCounter++
		}
		return availableLanguages["en"], nil
	}
	_, err := countries.New(languagesRepository, true)
	if err == nil {
		t.Fatal("error should not be nil")
	}
	if targetError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", targetError.Error(), err.Error())
	}
}
