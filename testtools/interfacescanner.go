package testtools

import (
	"database/sql"
	"time"

	"github.com/juju/errors"
)

// TODO :: 7777 These need more types. Can just add on-the-fly when tests need them,
// or maybe just add them all at once to prevent the hassle of understanding this
// over and over again.

var errNilPtr = errors.New("destination pointer is nil")

// ScanInterfaceValues is a test helper to resolve scanning anonymous
// values into each other.
func ScanInterfaceValues(src []interface{}, dest []interface{}) error {
	for k, value := range src {
		switch v := value.(type) {
		case string:
			*dest[k].(*string) = value.(string)
		case uint16:
			*dest[k].(*uint16) = value.(uint16)
		case bool:
			*dest[k].(*bool) = value.(bool)
		case time.Time:
			*dest[k].(*time.Time) = value.(time.Time)
		case nil:
			switch d := dest[k].(type) {
			case *interface{}:
				if d == nil {
					return errNilPtr
				}
				*d = nil
				return nil
			case *[]byte:
				if d == nil {
					return errNilPtr
				}
				*d = nil
				return nil
			case *sql.RawBytes:
				if d == nil {
					return errNilPtr
				}
				*d = nil
				return nil
			}
		default:
			return errors.Errorf("uhandled type %v %v", v, value)
		}
	}
	return nil
}
