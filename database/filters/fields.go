package filters

import (
	"strconv"

	"github.com/espal-digital-development/system/units"
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
func (f *field) UintValue() (uint, error) {
	if f.value == "" {
		return 0, nil
	}
	ui64, err := strconv.ParseUint(f.value, units.Base10, units.BitWidth64Bit)
	if err != nil {
		return 0, errors.Trace(err)
	}
	return uint(ui64), nil
}
