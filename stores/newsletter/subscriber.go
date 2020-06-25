package newsletter

import (
	"time"
)

// Subscriber database object.
// @synthesize
type Subscriber struct {
	// TODO :: Logic is `OR` User linked e-mail `OR` only Email without User
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	newsletterID       string
	userID             *string
	email              *string
}

// TableName returns the table name that belongs to the current model.
func (s *Subscriber) TableName() string {
	return "NewsletterSubscriber"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *Subscriber) TableAlias() string {
	return "ns"
}
