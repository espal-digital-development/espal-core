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
func (c *Currency) ID() uint {
	return c.id
}

// Code returns the internal Currency code.
func (c *Currency) Code() string {
	return c.code
}

// Numeral returns the internal Currency numeral code.
func (c *Currency) Numeral() string {
	return c.numeral
}

// Symbol returns the internal Currency symbol.
func (c *Currency) Symbol() string {
	return c.symbol
}

// EnglishName returns the internal Currency englishName.
func (c *Currency) EnglishName() string {
	return c.englishName
}

// SetTranslation sets a translations for the Language.
func (c *Currency) SetTranslation(languageID uint16, value string) {
	c.translations[languageID] = value
}

// HasTranslation returns an indicator if the given currency
// has a translation for the given locale.
func (c *Currency) HasTranslation(localeID uint16) (string, bool) {
	value, ok := c.translations[localeID]
	return value, ok
}

// Translate get's the Currency's translation based on the given locale.
func (c *Currency) Translate(localeID uint16) string {
	if localeID == c.englishLocaleID {
		return c.englishName
	}
	if _, ok := c.translations[localeID]; !ok {
		return c.englishName
	}
	return c.translations[localeID]
}
