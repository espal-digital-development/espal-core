package filters

// Context represents a filterable's object extra options.
type Context interface {
	Translate(string) string
	TranslatePlural(string) string
}
