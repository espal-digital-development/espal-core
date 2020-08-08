package queryhelper

import (
	"strconv"
	"strings"
)

var _ Helper = &QueryHelper{}

const paramFieldWithComma = 3

// Helper represents an object that can help build and mutate database queries.
type Helper interface {
	BuildUpdateWhereInIds(tableName string, operation string, fieldName string, values []string) (string,
		[]interface{}, error)
	BuildDeleteWhereInIds(tableName string, fieldName string, values []string) (string, []interface{}, error)
}

// QueryHelper is an assistant service to help build and mutate queries.
type QueryHelper struct{}

// BuildUpdateWhereInIds takes a table name and field name to match the given ids against to build a UDPDATE WHERE IN
// query.
func (h *QueryHelper) BuildUpdateWhereInIds(tableName string, operation string, fieldName string,
	ids []string) (string, []interface{}, error) {
	idsLength := len(ids)
	idsInterfaces := make([]interface{}, idsLength)
	for k := range ids {
		idsInterfaces[k] = ids[k]
	}
	whereInParams := &strings.Builder{}
	whereInParams.WriteString(`UPDATE SET `)
	whereInParams.WriteString(operation)
	whereInParams.WriteString(` "`)
	whereInParams.WriteString(tableName)
	whereInParams.WriteString(`" WHERE "`)
	whereInParams.WriteString(fieldName)
	whereInParams.WriteString(`" IN (`)
	whereInParams.Grow(idsLength*paramFieldWithComma - 1)
	for i := 1; i <= idsLength; i++ {
		whereInParams.WriteString("$")
		whereInParams.WriteString(strconv.Itoa(i))
	}
	whereInParams.WriteString(`)`)
	return whereInParams.String(), idsInterfaces, nil
}

// BuildDeleteWhereInIds takes a table name and field name to match the given ids against to build a DELETE WHERE IN
// query.
func (h *QueryHelper) BuildDeleteWhereInIds(tableName string, fieldName string,
	ids []string) (string, []interface{}, error) {
	idsLength := len(ids)
	idsInterfaces := make([]interface{}, idsLength)
	for k := range ids {
		idsInterfaces[k] = ids[k]
	}
	whereInParams := &strings.Builder{}
	whereInParams.WriteString(`DELETE FROM "`)
	whereInParams.WriteString(tableName)
	whereInParams.WriteString(`" WHERE "`)
	whereInParams.WriteString(fieldName)
	whereInParams.WriteString(`" IN (`)
	whereInParams.Grow(idsLength*paramFieldWithComma - 1)
	for i := 1; i <= idsLength; i++ {
		whereInParams.WriteString("$")
		whereInParams.WriteString(strconv.Itoa(i))
	}
	whereInParams.WriteString(`)`)
	return whereInParams.String(), idsInterfaces, nil
}

// New returns a new instance of QueryHelper.
func New() *QueryHelper {
	return &QueryHelper{}
}
