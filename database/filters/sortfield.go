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
func (f *sortField) Name() string {
	return f.name
}

// TableAlias returns the field it's table alias.
func (f *sortField) TableAlias() string {
	return f.tableAlias
}

// Descending returns if the field should be sorted descending.
func (f *sortField) Descending() bool {
	return f.descending
}

// SetTableAlias sets the field table alias.
func (f *sortField) SetTableAlias(tableAlias string) SortField {
	f.tableAlias = tableAlias
	return f
}

// NewSortField returns a new instance of a SortField.
func (f *filter) NewSortField(name string, descending bool) SortField {
	return &sortField{
		name:       name,
		descending: descending,
	}
}
