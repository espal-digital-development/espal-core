package countries

var _ Data = &Country{}

// Data represents an object that gives country data.
type Data interface {
	ID() uint16
	Code() string
	EnglishName() string
	SetTranslation(languageID uint16, value string)
	HasTranslation(localeID uint16) (string, bool)
	Translate(localeID uint16) string
}

// Country struct containing the Country's data.
type Country struct {
	id              uint16
	englishLocaleID uint16
	code            string
	englishName     string
	translations    map[uint16]string
}

// ID returns the internal Country id.
func (c *Country) ID() uint16 {
	return c.id
}

// Code returns the internal Country code.
func (c *Country) Code() string {
	return c.code
}

// EnglishName returns the internal Country englishName.
func (c *Country) EnglishName() string {
	return c.englishName
}

// SetTranslation sets a translations for the Country.
func (c *Country) SetTranslation(languageID uint16, value string) {
	c.translations[languageID] = value
}

// HasTranslation returns an indicator if the given country
// has a translation for the given locale.
func (c *Country) HasTranslation(localeID uint16) (string, bool) {
	value, ok := c.translations[localeID]
	return value, ok
}

// Translate get's the Country's translation based on the given locale.
func (c *Country) Translate(localeID uint16) string {
	if localeID == c.englishLocaleID {
		return c.englishName
	}
	if _, ok := c.translations[localeID]; !ok {
		return c.englishName
	}
	return c.translations[localeID]
}
