package site

import (
	"strings"
	"time"
)

// nolint:deadcode
type siteMethods interface {
	CurrenciesCount() uint
}

// Site database object.
// @synthesize
type Site struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	online             bool
	language           *uint16
	country            *uint16
	currencies         string

	localizedName *string // @synthesize-no-db-field
}

// CurrenciesCount returns the amount of currencies that are
// defined in the string field.
func (s *Site) CurrenciesCount() uint {
	if s.currencies != "" {
		return uint(strings.Count(s.currencies, ",") + 1)
	}
	return 0
}
