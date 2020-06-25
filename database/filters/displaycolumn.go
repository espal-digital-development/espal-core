package filters

// DisplayColumn holds database column name respresentations.
type DisplayColumn interface {
	Name() string
	Plural() bool
	SetPlural(plural bool) DisplayColumn
}

type displayColumn struct {
	name   string
	plural bool
}

// Name returns the column name.
func (c *displayColumn) Name() string {
	return c.name
}

// Name returns the if the column name is plural.
func (c *displayColumn) Plural() bool {
	return c.plural
}

// SetPlural sets the column plural.
func (c *displayColumn) SetPlural(plural bool) DisplayColumn {
	c.plural = plural
	return c
}

// NewColumn returns a new instance of a DisplayColumn.
func (f *filter) NewColumn(name string) DisplayColumn {
	return &displayColumn{
		name: name,
	}
}
