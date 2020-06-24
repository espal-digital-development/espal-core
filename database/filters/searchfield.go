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
func (searchField *searchField) Name() string {
	return searchField.name
}

// TableAlias returns the field it's table alias.
func (searchField *searchField) TableAlias() string {
	return searchField.tableAlias
}

// SetTableAlias sets the field table alias.
func (searchField *searchField) SetTableAlias(tableAlias string) SearchField {
	searchField.tableAlias = tableAlias
	return searchField
}

// NewSearchField returns a new instance of a SearchField.
func (filter *filter) NewSearchField(name string) SearchField {
	return &searchField{
		name: name,
	}
}
