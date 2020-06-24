package saleoffer

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: PreSave/PreUpdate only allowed Variant OR Bundled OR none (custom line)
	// TODO :: Discount amounts (percentual/fixed) should be separate lines here?
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	sorting            uint
	saleOfferID        string
	productVariantID   *string
	bundledProductID   *string
	quantity           int
	price              float32 // Price per unit
	vatPercentage      float32
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (line *Line) TableName() string {
	return "SaleOfferLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (line *Line) TableAlias() string {
	return "sofl"
}
