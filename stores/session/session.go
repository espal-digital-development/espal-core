package session

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/juju/errors"
)

// nolint:deadcode
type sessionMethods interface {
	SetDataFromJSON(entries DataEntries) error
	GetDataAsJSON() (DataEntries, error)
}

// Session database object.
// @synthesize
type Session struct {
	id                 string
	createdByID        *string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	timeout            time.Duration
	hash               string
	data               []byte

	dataEntries *dataEntries // @synthesize-no-db-field
}

// TableAlias returns the unique resolved table alias for use in queries.
func (session *Session) TableAlias() string {
	return "ses"
}

// SetDataFromJSON sets the JSON data as bytes into the data.
func (session *Session) SetDataFromJSON(entries DataEntries) error {
	session.dataEntries = newDataEntries()
	for key, value := range entries.All() {
		session.dataEntries.entries[key] = value
	}
	var err error
	session.data, err = jsoniter.Marshal(session.dataEntries.entries)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// GetDataAsJSON returns the marshalled JSON of the data bytes.
func (session *Session) GetDataAsJSON() (DataEntries, error) {
	if len(session.data) == 0 {
		return nil, nil
	}
	session.dataEntries = newDataEntries()
	if err := jsoniter.Unmarshal(session.data, &session.dataEntries.entries); err != nil {
		return nil, errors.Trace(err)
	}
	return session.dataEntries, nil
}

func newSession() *Session {
	return &Session{
		dataEntries: newDataEntries(),
	}
}
