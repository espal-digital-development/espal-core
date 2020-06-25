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
func (s *Session) TableAlias() string {
	return "ses"
}

// SetDataFromJSON sets the JSON data as bytes into the data.
func (s *Session) SetDataFromJSON(entries DataEntries) error {
	s.dataEntries = newDataEntries()
	for key, value := range entries.All() {
		s.dataEntries.entries[key] = value
	}
	var err error
	s.data, err = jsoniter.Marshal(s.dataEntries.entries)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// GetDataAsJSON returns the marshalled JSON of the data bytes.
func (s *Session) GetDataAsJSON() (DataEntries, error) {
	if len(s.data) == 0 {
		return nil, nil
	}
	s.dataEntries = newDataEntries()
	if err := jsoniter.Unmarshal(s.data, &s.dataEntries.entries); err != nil {
		return nil, errors.Trace(err)
	}
	return s.dataEntries, nil
}

func newSession() *Session {
	return &Session{
		dataEntries: newDataEntries(),
	}
}
