package group

import (
	"strings"
	"time"
)

// nolint:deadcode
type groupMethods interface {
	UserRightsCount() uint
	CurrenciesCount() uint
}

// Group database object.
// @synthesize
type Group struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	active             bool
	userRights         string
	currencies         string

	localizedName *string // @synthesize-no-db-field
}

// TableName returns the table name that belongs to the current model.
func (g *Group) TableName() string {
	return "UserGroup"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (g *Group) TableAlias() string {
	return "ug"
}

// UserRightsCount returns the amount of userRights that are
// defined in the string field.
func (g *Group) UserRightsCount() uint {
	if g.userRights != "" {
		return uint(strings.Count(g.userRights, ",") + 1)
	}
	return 0
}

// CurrenciesCount returns the amount of currencies that are
// defined in the string field.
func (g *Group) CurrenciesCount() uint {
	if g.currencies != "" {
		return uint(strings.Count(g.currencies, ",") + 1)
	}
	return 0
}
