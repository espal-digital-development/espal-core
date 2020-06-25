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
func (l *Language) ID() uint16 {
	return l.id
}

// Code returns the internalanguage Language code.
func (l *Language) Code() string {
	return l.code
}

// EnglishName returns the internalanguage Language englishName.
func (l *Language) EnglishName() string {
	return l.englishName
}

// SetTranslation sets a translations for the l.
func (l *Language) SetTranslation(languageID uint16, value string) {
	l.translations[languageID] = value
}

// AlternativeEnglishName returns the internalanguage Language alternativeEnglishName.
func (l *Language) AlternativeEnglishName() string {
	return l.alternativeEnglishName
}

// Translate get's the Language's translation based on the given locale.
func (l *Language) Translate(localeID uint16) string {
	if localeID == l.englishLocaleID {
		return l.EnglishName()
	}
	if _, ok := l.translations[localeID]; !ok {
		return l.EnglishName()
	}
	return l.translations[localeID]
}
