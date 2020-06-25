package price

import (
	"time"
)

// Price database object.
// @synthesize
type Price struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	bundledProductID   *string
	productModelID     *string
	productVariantID   *string
	domainID           string
	priceGroupID       string
	taxGroup           uint
	currency           uint16
	price              float32
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *Price) TableAlias() string {
	return "pr"
}
