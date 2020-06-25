package product

// PropertyRevisionHistory database object.
// @synthesize
type PropertyRevisionHistory struct {
	PropertyRevision
}

// TableName returns the table name that belongs to the current model.
func (p *PropertyRevisionHistory) TableName() string {
	return "ProductPropertyRevisionHistory"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *PropertyRevisionHistory) TableAlias() string {
	return "proprh"
}
