package price

import (
	"time"
)

// MinMax database object.
// @synthesize
type MinMax struct {
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
	lowest             float64
	highest            float64
	lowestRev          float64 // This field is persistently auto-filled
	highestRev         float64 // This field is persistently auto-filled
}

// TableName returns the table name that belongs to the current model.
func (minMax *MinMax) TableName() string {
	return "PriceMinMax"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (minMax *MinMax) TableAlias() string {
	return "prmm"
}
