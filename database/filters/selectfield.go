package filters

// SelectField holds database filter field selection.
type SelectField interface {
	Name() string
	Alias() string
	SetAlias(alias string) SelectField
	SetMapTo(mapTo string) SelectField
}

type selectField struct {
	name  string
	alias string
	mapTo string
}

// Name returns the field name.
func (f *selectField) Name() string {
	return f.name
}

// TableAlias returns the field it's alias.
func (f *selectField) Alias() string {
	return f.alias
}

// SetAlias sets the field alias.
func (f *selectField) SetAlias(alias string) SelectField {
	f.alias = alias
	return f
}

// SetMapTo sets the field mapTo.
func (f *selectField) SetMapTo(mapTo string) SelectField {
	f.mapTo = mapTo
	return f
}

// NewSelectField returns a new instance of a SelectField.
func (f *filter) NewSelectField(name string) SelectField {
	return &selectField{
		name: name,
	}
}
