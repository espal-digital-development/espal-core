package filters

import (
	"strconv"

	"github.com/juju/errors"
)

type fieldAction uint8

// FilterFieldAction for database comparison actions.
const (
	filterFieldActionIsNull fieldAction = iota + 1
	filterFieldActionEqualTo
	filterFieldActionNotEqualTo
	filterFieldActionGreaterThan
	filterFieldActionSmallerThan
	filterFieldActionEqualToOrGreaterThan
	filterFieldActionEqualToOrSmallerThan
	filterFieldActionBetween
	filterFieldActionLike
)

type field struct {
	name   string
	value  string
	value2 string // Used for range-inspection filters
	action fieldAction
	alias  string
}

// UintValue returns the filter field value as uint.
func (field *field) UintValue() (uint, error) {
	if field.value == "" {
		return 0, nil
	}
	ui64, err := strconv.ParseUint(field.value, 10, 64)
	if err != nil {
		return 0, errors.Trace(err)
	}
	return uint(ui64), nil
}
