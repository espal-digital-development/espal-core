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
func (group *Group) TableName() string {
	return "UserGroup"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (group *Group) TableAlias() string {
	return "ug"
}

// UserRightsCount returns the amount of userRights that are
// defined in the string field.
func (group *Group) UserRightsCount() uint {
	if group.userRights != "" {
		return uint(strings.Count(group.userRights, ",") + 1)
	}
	return 0
}

// CurrenciesCount returns the amount of currencies that are
// defined in the string field.
func (group *Group) CurrenciesCount() uint {
	if group.currencies != "" {
		return uint(strings.Count(group.currencies, ",") + 1)
	}
	return 0
}
