package translations

// Config Translations provider object.
type Config struct {
	Path string
}

// Translations provider object.
type Translations struct {
	path string
}

// New returns a new instance of Translations.
func New(config *Config) (*Translations, error) {
	t := &Translations{
		path: config.Path,
	}
	return t, nil
}
