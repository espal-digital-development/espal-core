package languages_test

import (
	"strings"
	"testing"

	"github.com/espal-digital-development/espal-core/repositories/languages"
)

func getDefault(t *testing.T) languages.Repository {
	languages, err := languages.New()
	if err != nil {
		t.Fatal(err)
	}
	return languages
}

func TestAll(t *testing.T) {
	languages := getDefault(t)

	all := languages.All()
	if len(all) != 613 {
		t.Fatal("All returns an incorrect amount")
	}

	for id := range all {
		l, err := languages.ByID(id)
		if err != nil {
			t.Fatal(err)
		}
		code := l.Code()
		if strings.Trim(code, " ") == "" {
			t.Errorf("`%s` is an invalid code", code)
		}
		name := l.EnglishName()
		if len(name) < 2 {
			t.Errorf("name of `%s` shorter than 2 characters", name)
		}
	}
}

func TestByID(t *testing.T) {
	languages := getDefault(t)

	l, err := languages.ByID(141)
	if err != nil {
		t.Fatal(err)
	}
	if l.ID() != 141 {
		t.Errorf("expected ID `141` but got `%d`", l.ID())
	}

	_, err = languages.ByID(0)
	if err == nil {
		t.Error("calling unexisting should give an error")
	}
}

func TestByCode(t *testing.T) {
	languages := getDefault(t)

	c, err := languages.ByCode("en")
	if err != nil {
		t.Fatal(err)
	}
	if c.Code() != "en" {
		t.Errorf("expected code `en` but got `%s`", c.Code())
	}

	_, err = languages.ByCode("00")
	if err == nil {
		t.Error("calling unexisting Country should give an error")
	}
}

func TestLanguageID(t *testing.T) {
	languages := getDefault(t)

	language, err := languages.ByCode("en")
	if err != nil {
		t.Fatal(err)
	}
	if language.ID() == 0 {
		t.Fatal("ID should not be 0")
	}
}

func TestLanguageAlternativeEnglishName(t *testing.T) {
	languages := getDefault(t)

	language, err := languages.ByCode("en")
	if err != nil {
		t.Fatal(err)
	}
	language.AlternativeEnglishName()
}

func TestLanguageTranslate(t *testing.T) {
	languages := getDefault(t)

	l, err := languages.ByCode("es")
	if err != nil {
		t.Fatal(err)
	}
	esTrans := l.Translate(l.ID())
	if esTrans != "español" {
		t.Errorf("translation for Spanish should return español, but returned `%s`", esTrans)
	}
	l2, err := languages.ByCode("en")
	if err != nil {
		t.Fatal(err)
	}
	l.Translate(l2.ID())
	l.Translate(0)
}

func TestEnglishLocaleID(t *testing.T) {
	languages := getDefault(t)
	if languages.EnglishLocaleID() != 141 {
		t.Fatal("locale ID for English should be 141")
	}
}
