package filters

// SortField holds filtering mechanics options.
type SortField interface {
	Name() string
	TableAlias() string
	Descending() bool
	SetTableAlias(tableAlias string) SortField
}

type sortField struct {
	name       string
	tableAlias string
	descending bool
}

// Name returns the field name.
func (sortField *sortField) Name() string {
	return sortField.name
}

// TableAlias returns the field it's table alias.
func (sortField *sortField) TableAlias() string {
	return sortField.tableAlias
}

// Descending returns if the field should be sorted descending.
func (sortField *sortField) Descending() bool {
	return sortField.descending
}

// SetTableAlias sets the field table alias.
func (sortField *sortField) SetTableAlias(tableAlias string) SortField {
	sortField.tableAlias = tableAlias
	return sortField
}

// NewSortField returns a new instance of a SortField.
func (filter *filter) NewSortField(name string, descending bool) SortField {
	return &sortField{
		name:       name,
		descending: descending,
	}
}
