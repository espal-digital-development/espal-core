package translations_test

import (
	"math"
	"testing"

	"github.com/espal-digital-development/espal-core/logger/loggermock"
	"github.com/espal-digital-development/espal-core/repositories/translations"
)

const (
	stubKey       = "key"
	stubExpected1 = "Expected"
	stubExpected2 = "Expected2"
)

var (
	stubLanguage1ID    uint16
	stubLanguage2ID    uint16
	availableLanguages map[uint16]string
	loggerService      *loggermock.LoggerMock
)

func initMocks() {
	stubLanguage1ID = math.MaxUint16 - 1
	stubLanguage2ID = math.MaxUint16 - 2
	availableLanguages = map[uint16]string{
		stubLanguage1ID: "xyz1",
		stubLanguage2ID: "xyz2",
	}
	loggerService = &loggermock.LoggerMock{
		GetLastMessageFunc: func() string {
			return ""
		},
		WarningfFunc: func(message string, params ...interface{}) {
		},
	}
}

func getDefault(t *testing.T) *translations.Translations {
	initMocks()
	translations, err := translations.New(loggerService, availableLanguages)
	if err != nil {
		t.Fatal(err)
	}
	return translations
}

func TestAvailableLanguage(t *testing.T) {
	var englishLanguageID uint16 = 141
	languages := map[uint16]string{englishLanguageID: "en"}
	_, err := translations.New(loggerService, languages)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnexistingLanguage(t *testing.T) {
	translations := getDefault(t)
	translations.Singular(1001, "user")
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing language")
	}
}

func TestCheckExistenceNotExisting(t *testing.T) {
	initMocks()
	var englishLanguageID uint16 = 141
	languages := map[uint16]string{englishLanguageID: "en"}
	translations, err := translations.New(loggerService, languages)
	if err != nil {
		t.Fatal(err)
	}
	translations.Singular(englishLanguageID, stubKey)
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestLoadForLanguageFromYamlData(t *testing.T) {
	translations := getDefault(t)
	yamlData := []byte(`someValue:
  s: "Some value"
`)
	if err := translations.LoadForLanguageFromYamlData(stubLanguage1ID, yamlData); err != nil {
		t.Fatal(err)
	}
}

func TestLoadForLanguageFromYamlDataExisting(t *testing.T) {
	translations := getDefault(t)
	stubKey2 := "key2"
	if err := translations.SetSingular(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	if err := translations.SetFormattedPlural(stubLanguage1ID, stubKey2, stubExpected2); err != nil {
		t.Fatal(err)
	}
	yamlData := []byte(stubKey + `:
  p: "Some values"
  sf: "Some value %s"
  pf: "Some values %s"
` + stubKey2 + `:
  s: "Some value"
  p: "Some values"
  sf: "Some value %s"
`)
	if err := translations.LoadForLanguageFromYamlData(stubLanguage1ID, yamlData); err != nil {
		t.Fatal(err)
	}
}

func TestLoadForLanguageFromYamlDataUnmarshalError(t *testing.T) {
	translations := getDefault(t)
	yamlData := []byte(`broken%yaml`)
	if err := translations.LoadForLanguageFromYamlData(stubLanguage1ID, yamlData); err == nil {
		t.Fatal("broken yaml data should give an error")
	}
}

func TestSingular(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetSingular(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	f := translations.Singular(stubLanguage1ID, stubKey)
	if stubExpected1 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected1, f)
	}
	if err := translations.SetSingular(stubLanguage2ID, stubKey, stubExpected2); err != nil {
		t.Fatal(err)
	}
	f = translations.Singular(stubLanguage2ID, stubKey)
	if stubExpected2 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected2, f)
	}
}

func TestSingularExists(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetSingular(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	expectedErr := "singular translation already set for key `" + stubKey + "`"
	err := translations.SetSingular(stubLanguage1ID, stubKey, stubExpected1)
	if err == nil {
		t.Fatal("setting singular twice should give an error")
	}
	if err.Error() != expectedErr {
		t.Fatalf("setting singular twice should give error `%s`, but got `%s`", expectedErr, err.Error())
	}
}

func TestSingularDoesntExist(t *testing.T) {
	translations := getDefault(t)
	f := translations.Singular(stubLanguage1ID, stubKey)
	if stubKey != f {
		t.Fatalf("Expected `%s` but got `%s`", stubKey, f)
	}
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestSingularTypeDoesntExist(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetPlural(stubLanguage1ID, stubKey, "stub"); err != nil {
		t.Fatal(err)
	}
	translations.Singular(stubLanguage1ID, stubKey)
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestPlural(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetPlural(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	f := translations.Plural(stubLanguage1ID, stubKey)
	if stubExpected1 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected1, f)
	}
	if err := translations.SetPlural(stubLanguage2ID, stubKey, stubExpected2); err != nil {
		t.Fatal(err)
	}
	f = translations.Plural(stubLanguage2ID, stubKey)
	if stubExpected2 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected2, f)
	}
}

func TestPluralExists(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetPlural(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	expectedErr := "plural translation already set for key `" + stubKey + "`"
	err := translations.SetPlural(stubLanguage1ID, stubKey, stubExpected1)
	if err == nil {
		t.Fatal("setting plural twice should give an error")
	}
	if err.Error() != expectedErr {
		t.Fatalf("setting plural twice should give error `%s`, but got `%s`", expectedErr, err.Error())
	}
}

func TestPluralDoesntExist(t *testing.T) {
	translations := getDefault(t)
	f := translations.Plural(stubLanguage1ID, stubKey)
	if stubKey != f {
		t.Fatalf("Expected `%s` but got `%s`", stubKey, f)
	}
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestPluralTypeDoesntExist(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetSingular(stubLanguage1ID, stubKey, "stub"); err != nil {
		t.Fatal(err)
	}
	translations.Plural(stubLanguage1ID, stubKey)
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestFormatted(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetFormatted(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	f := translations.Formatted(stubLanguage1ID, stubKey)
	if stubExpected1 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected1, f)
	}
	if err := translations.SetFormatted(stubLanguage2ID, stubKey, stubExpected2); err != nil {
		t.Fatal(err)
	}
	f = translations.Formatted(stubLanguage2ID, stubKey)
	if stubExpected2 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected2, f)
	}
}

func TestFormattedExists(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetFormatted(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	expectedErr := "formatted translation already set for key `" + stubKey + "`"
	err := translations.SetFormatted(stubLanguage1ID, stubKey, stubExpected1)
	if err == nil {
		t.Fatal("setting formatted twice should give an error")
	}
	if err.Error() != expectedErr {
		t.Fatalf("setting formatted twice should give error `%s`, but got `%s`", expectedErr, err.Error())
	}
}

func TestFormattedDoesntExist(t *testing.T) {
	translations := getDefault(t)
	f := translations.Formatted(stubLanguage1ID, stubKey)
	if stubKey != f {
		t.Fatalf("Expected `%s` but got `%s`", stubKey, f)
	}
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestFormattedTypeDoesntExist(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetSingular(stubLanguage1ID, stubKey, "stub"); err != nil {
		t.Fatal(err)
	}
	translations.Formatted(stubLanguage1ID, stubKey)
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestFormattedPlural(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetFormattedPlural(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	f := translations.FormattedPlural(stubLanguage1ID, stubKey)
	if stubExpected1 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected1, f)
	}
	if err := translations.SetFormattedPlural(stubLanguage2ID, stubKey, stubExpected2); err != nil {
		t.Fatal(err)
	}
	f = translations.FormattedPlural(stubLanguage2ID, stubKey)
	if stubExpected2 != f {
		t.Fatalf("Expected `%s` but got `%s`", stubExpected2, f)
	}
}

func TestFormattedPluralExists(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetFormattedPlural(stubLanguage1ID, stubKey, stubExpected1); err != nil {
		t.Fatal(err)
	}
	expectedErr := "pluralFormatted translation already set for key `" + stubKey + "`"
	err := translations.SetFormattedPlural(stubLanguage1ID, stubKey, stubExpected1)
	if err == nil {
		t.Fatal("setting formattedPlural twice should give an error")
	}
	if err.Error() != expectedErr {
		t.Fatalf("setting formattedPlural twice should give error `%s`, but got `%s`", expectedErr, err.Error())
	}
}

func TestFormattedPluralDoesntExist(t *testing.T) {
	translations := getDefault(t)
	f := translations.FormattedPlural(stubLanguage1ID, stubKey)
	if stubKey != f {
		t.Fatalf("Expected `%s` but got `%s`", stubKey, f)
	}
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}

func TestFormattedPluralTypeDoesntExist(t *testing.T) {
	translations := getDefault(t)
	if err := translations.SetSingular(stubLanguage1ID, stubKey, "stub"); err != nil {
		t.Fatal(err)
	}
	translations.FormattedPlural(stubLanguage1ID, stubKey)
	if len(loggerService.WarningfCalls()) == 0 {
		t.Fatal("logger should've been triggered for the missing translation")
	}
}
