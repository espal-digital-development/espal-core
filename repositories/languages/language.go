package languages

var _ Data = &Language{}

// Data represents an object that gives language data.
type Data interface {
	ID() uint16
	Code() string
	EnglishName() string
	SetTranslation(languageID uint16, value string)
	AlternativeEnglishName() string
	Translate(localeID uint16) string
}

// Language struct containing the Language's data.
type Language struct {
	id                     uint16
	englishLocaleID        uint16
	code                   string
	englishName            string
	alternativeEnglishName string
	translations           map[uint16]string
}

// ID returns the internalanguage Language id.
func (language *Language) ID() uint16 {
	return language.id
}

// Code returns the internalanguage Language code.
func (language *Language) Code() string {
	return language.code
}

// EnglishName returns the internalanguage Language englishName.
func (language *Language) EnglishName() string {
	return language.englishName
}

// SetTranslation sets a translations for the Language.
func (language *Language) SetTranslation(languageID uint16, value string) {
	language.translations[languageID] = value
}

// AlternativeEnglishName returns the internalanguage Language alternativeEnglishName.
func (language *Language) AlternativeEnglishName() string {
	return language.alternativeEnglishName
}

// Translate get's the Language's translation based on the given locale.
func (language *Language) Translate(localeID uint16) string {
	if localeID == language.englishLocaleID {
		return language.EnglishName()
	}
	if _, ok := language.translations[localeID]; !ok {
		return language.EnglishName()
	}
	return language.translations[localeID]
}
