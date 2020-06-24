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
func (displayColumn *displayColumn) Name() string {
	return displayColumn.name
}

// Name returns the if the column name is plural.
func (displayColumn *displayColumn) Plural() bool {
	return displayColumn.plural
}

// SetPlural sets the column plural.
func (displayColumn *displayColumn) SetPlural(plural bool) DisplayColumn {
	displayColumn.plural = plural
	return displayColumn
}

// NewColumn returns a new instance of a DisplayColumn.
func (filter *filter) NewColumn(name string) DisplayColumn {
	return &displayColumn{
		name: name,
	}
}
