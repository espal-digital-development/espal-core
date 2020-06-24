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
func (country *Country) ID() uint16 {
	return country.id
}

// Code returns the internal Country code.
func (country *Country) Code() string {
	return country.code
}

// EnglishName returns the internal Country englishName.
func (country *Country) EnglishName() string {
	return country.englishName
}

// SetTranslation sets a translations for the Country.
func (country *Country) SetTranslation(languageID uint16, value string) {
	country.translations[languageID] = value
}

// HasTranslation returns an indicator if the given country
// has a translation for the given locale.
func (country *Country) HasTranslation(localeID uint16) (string, bool) {
	value, ok := country.translations[localeID]
	return value, ok
}

// Translate get's the Country's translation based on the given locale.
func (country *Country) Translate(localeID uint16) string {
	if localeID == country.englishLocaleID {
		return country.englishName
	}
	if _, ok := country.translations[localeID]; !ok {
		return country.englishName
	}
	return country.translations[localeID]
}
