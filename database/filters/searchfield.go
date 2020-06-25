package filters

// SearchField holds database filter search fields.
type SearchField interface {
	Name() string
	TableAlias() string
	SetTableAlias(tableAlias string) SearchField
}

type searchField struct {
	name       string
	tableAlias string
}

// Name returns the field name.
func (f *searchField) Name() string {
	return f.name
}

// TableAlias returns the field it's table alias.
func (f *searchField) TableAlias() string {
	return f.tableAlias
}

// SetTableAlias sets the field table alias.
func (f *searchField) SetTableAlias(tableAlias string) SearchField {
	f.tableAlias = tableAlias
	return f
}

// NewSearchField returns a new instance of a SearchField.
func (f *filter) NewSearchField(name string) SearchField {
	return &searchField{
		name: name,
	}
}
