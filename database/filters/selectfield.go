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
func (selectField *selectField) Name() string {
	return selectField.name
}

// TableAlias returns the field it's alias.
func (selectField *selectField) Alias() string {
	return selectField.alias
}

// SetAlias sets the field alias.
func (selectField *selectField) SetAlias(alias string) SelectField {
	selectField.alias = alias
	return selectField
}

// SetMapTo sets the field mapTo.
func (selectField *selectField) SetMapTo(mapTo string) SelectField {
	selectField.mapTo = mapTo
	return selectField
}

// NewSelectField returns a new instance of a SelectField.
func (filter *filter) NewSelectField(name string) SelectField {
	return &selectField{
		name: name,
	}
}
