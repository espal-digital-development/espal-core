package currencies

var _ Data = &Currency{}

// Data represents an object that gives currency data.
type Data interface {
	ID() uint
	Code() string
	Symbol() string
	EnglishName() string
	SetTranslation(languageID uint16, value string)
	HasTranslation(localeID uint16) (string, bool)
	Translate(localeID uint16) string
}

// Currency struct containing the Currency's data.
type Currency struct {
	id              uint
	code            string
	numeral         string
	symbol          string
	englishName     string
	translations    map[uint16]string
	englishLocaleID uint16
}

// ID returns the internal Currency id.
func (currency *Currency) ID() uint {
	return currency.id
}

// Code returns the internal Currency code.
func (currency *Currency) Code() string {
	return currency.code
}

// Numeral returns the internal Currency numeral code.
func (currency *Currency) Numeral() string {
	return currency.numeral
}

// Symbol returns the internal Currency symbol.
func (currency *Currency) Symbol() string {
	return currency.symbol
}

// EnglishName returns the internal Currency englishName.
func (currency *Currency) EnglishName() string {
	return currency.englishName
}

// SetTranslation sets a translations for the Language.
func (currency *Currency) SetTranslation(languageID uint16, value string) {
	currency.translations[languageID] = value
}

// HasTranslation returns an indicator if the given currency
// has a translation for the given locale.
func (currency *Currency) HasTranslation(localeID uint16) (string, bool) {
	value, ok := currency.translations[localeID]
	return value, ok
}

// Translate get's the Currency's translation based on the given locale.
func (currency *Currency) Translate(localeID uint16) string {
	if localeID == currency.englishLocaleID {
		return currency.englishName
	}
	if _, ok := currency.translations[localeID]; !ok {
		return currency.englishName
	}
	return currency.translations[localeID]
}
