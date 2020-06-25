package emailtemplate

import (
	"time"
)

// Attachment database object.
// @synthesize
type Attachment struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	emailTemplateID    string
	filePath           string
	// Optional; if set to say `de` it would mean the attachment belongs to
	// the German translated version of the EmailTemplate.
	// TODO :: File-type and file header-/meta-info fields?
	language *uint16
}

// TableName returns the table name that belongs to the current model.
func (a *Attachment) TableName() string {
	return "EmailTemplateAttachment"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (a *Attachment) TableAlias() string {
	return "eta"
}
