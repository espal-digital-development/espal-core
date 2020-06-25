package group

import (
	"time"
)

// Group database object.
// @synthesize
type Group struct {
	// TODO :: M2M relational entity for PriceGroupsUsers and/or PriceGroupsUserGroups
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	code               string
	priority           uint // In case prices overlap; the highest priority is leading

	// Prices []Price
}

// TableName returns the table name that belongs to the current model.
func (g *Group) TableName() string {
	return "PriceGroup"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (g *Group) TableAlias() string {
	return "pg"
}
