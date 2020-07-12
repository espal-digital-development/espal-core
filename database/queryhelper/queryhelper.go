package queryhelper

import (
	"strconv"
	"strings"

	"github.com/juju/errors"
)

var _ Helper = &QueryHelper{}

const paramFieldWithComma = 3

// Helper represents an object that can help build and mutate database queries.
type Helper interface {
	BuildDeleteWhereInIds(tableName string, fieldName string, values []string) (string, []interface{}, error)
}

// QueryHelper is an assistant service to help build and mutate queries.
type QueryHelper struct{}

// BuildDeleteWhereInIds takes a table name and field name to match the given
// ids against to build a DELETE WHERE IN query.
func (h *QueryHelper) BuildDeleteWhereInIds(tableName string, fieldName string,
	ids []string) (string, []interface{}, error) {
	idsLength := len(ids)
	idsInterfaces := make([]interface{}, idsLength)
	for k := range ids {
		idsInterfaces[k] = ids[k]
	}
	whereInParams := &strings.Builder{}
	if _, err := whereInParams.WriteString(`DELETE FROM "`); err != nil {
		return "", nil, errors.Trace(err)
	}
	if _, err := whereInParams.WriteString(tableName); err != nil {
		return "", nil, errors.Trace(err)
	}
	if _, err := whereInParams.WriteString(`" WHERE "`); err != nil {
		return "", nil, errors.Trace(err)
	}
	if _, err := whereInParams.WriteString(fieldName); err != nil {
		return "", nil, errors.Trace(err)
	}
	if _, err := whereInParams.WriteString(`" IN (`); err != nil {
		return "", nil, errors.Trace(err)
	}
	whereInParams.Grow(idsLength*paramFieldWithComma - 1)
	for i := 1; i <= idsLength; i++ {
		if _, err := whereInParams.WriteString("$"); err != nil {
			return "", nil, errors.Trace(err)
		}
		if _, err := whereInParams.WriteString(strconv.Itoa(i)); err != nil {
			return "", nil, errors.Trace(err)
		}
	}
	if _, err := whereInParams.WriteString(`)`); err != nil {
		return "", nil, errors.Trace(err)
	}
	return whereInParams.String(), idsInterfaces, nil
}

// New returns a new instance of QueryHelper.
func New() *QueryHelper {
	return &QueryHelper{}
}
