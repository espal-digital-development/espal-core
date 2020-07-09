package translations

type Translations struct {
}

// New returns a new instance of Translations.
func New() (*Translations, error) {
	t := &Translations{}
	return t, nil
}
