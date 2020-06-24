package translations_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/logger/loggermock"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/storage/storagemock"
	"github.com/juju/errors"
)

var (
	englishLanguageID  uint16
	dutchLanguageID    uint16
	availableLanguages map[uint16]string
	files              map[string][]byte
	loggerService      *loggermock.LoggerMock
	storage            *storagemock.StorageMock
)

func initMocks() {
	englishLanguageID = 141
	dutchLanguageID = 388
	availableLanguages = map[uint16]string{
		englishLanguageID: "en",
		dutchLanguageID:   "nl",
	}
	files = map[string][]byte{
		"en.yml": []byte(`
user:
  s: User
domain:
  p: Domains
`),
		"nl.yml": []byte(`
user:
  s: Gebruiker
domain:
  p: Domeinen
`),
	}
	loggerService = &loggermock.LoggerMock{
		GetLastMessageFunc: func() string {
			return ""
		},
		WarningfFunc: func(message string, params ...interface{}) {
		},
	}
	storage = &storagemock.StorageMock{
		GetFunc: func(key string) ([]byte, bool, error) {
			if b, ok := files[key]; ok {
				return b, true, nil
			}
			return nil, false, errors.New("data not found")
		},
	}
}

func getDefault(t *testing.T) translations.Repository {
	initMocks()
	translations, err := translations.New(loggerService, storage, availableLanguages)
	if err != nil {
		t.Fatal(err)
	}
	return translations
}

func TestUnexistingLanguage(t *testing.T) {
	translations := getDefault(t)
	translations.Singular(1001, "user")
	if len(loggerService.WarningfCalls()) != 1 {
		t.Fatal("logger should've been triggered for the missing language")
	}
}

func TestGetFileError(t *testing.T) {
	initMocks()
	notFoundError := errors.New("data not found")
	storage.GetFunc = func(key string) ([]byte, bool, error) {
		return nil, false, notFoundError
	}
	_, err := translations.New(loggerService, storage, availableLanguages)
	if err == nil {
		t.Fatal("Error should not be nil")
	}
	if notFoundError != errors.Cause(err) {
		t.Fatalf("Expected `%s` but got `%s`", notFoundError.Error(), err.Error())
	}
}

func TestGetFileNotOk(t *testing.T) {
	initMocks()
	storage.GetFunc = func(key string) ([]byte, bool, error) {
		return nil, false, nil
	}
	_, err := translations.New(loggerService, storage, availableLanguages)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingular(t *testing.T) {
	translations := getDefault(t)
	expected := "User"
	f := translations.Singular(englishLanguageID, "user")
	if expected != f {
		t.Fatalf("Expected `%s` but got `%s`", expected, f)
	}
	expected = "Gebruiker"
	f = translations.Singular(dutchLanguageID, "user")
	if expected != f {
		t.Fatalf("Expected `%s` but got `%s`", expected, f)
	}
}

func TestSingularDoesntExist(t *testing.T) {
	translations := getDefault(t)
	expected := "idontexist"
	f := translations.Singular(englishLanguageID, expected)
	if expected != f {
		t.Fatalf("Expected `%s` but got `%s`", expected, f)
	}
	if len(loggerService.WarningfCalls()) != 1 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestSingularTypeDoesntExist(t *testing.T) {
	translations := getDefault(t)
	translations.Singular(englishLanguageID, "domain")
	if len(loggerService.WarningfCalls()) != 1 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestPlural(t *testing.T) {
	translations := getDefault(t)
	expected := "Domains"
	f := translations.Plural(englishLanguageID, "domain")
	if expected != f {
		t.Fatalf("Expected `%s` but got `%s`", expected, f)
	}
	expected = "Domeinen"
	f = translations.Plural(dutchLanguageID, "domain")
	if expected != f {
		t.Fatalf("Expected `%s` but got `%s`", expected, f)
	}
}

func TestPluralDoesntExist(t *testing.T) {
	translations := getDefault(t)
	expected := "idontexist"
	f := translations.Plural(englishLanguageID, expected)
	if expected != f {
		t.Fatalf("Expected `%s` but got `%s`", expected, f)
	}
	if len(loggerService.WarningfCalls()) != 1 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestPluralTypeDoesntExist(t *testing.T) {
	translations := getDefault(t)
	translations.Plural(englishLanguageID, "user")
	if len(loggerService.WarningfCalls()) != 1 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}
