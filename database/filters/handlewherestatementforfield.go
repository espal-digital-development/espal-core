package filters

import (
	"strconv"

	"github.com/juju/errors"
)

// nolint:gocyclo,funlen
func (f *filter) handleWhereStatementForField(fieldKey int) error {
	if f.firstWhereStatementHad {
		if _, err := f.query.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}
	} else {
		if _, err := f.query.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		f.firstWhereStatementHad = true
	}

	if f.fields[fieldKey].alias == "" {
		if _, err := f.query.WriteString(f.tableAlias); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString("."); err != nil {
			return errors.Trace(err)
		}
	} else if f.fields[fieldKey].alias != "" {
		if _, err := f.query.WriteString(f.fields[fieldKey].alias); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString("."); err != nil {
			return errors.Trace(err)
		}
	}

	if _, err := f.query.WriteString(f.fields[fieldKey].name); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.countQuery.WriteString(f.fields[fieldKey].name); err != nil {
		return errors.Trace(err)
	}

	var dontBindParam bool

	switch f.fields[fieldKey].action {
	case filterFieldActionEqualTo:
		if _, err := f.query.WriteString("="); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString("="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionIsNull:
		if f.fields[fieldKey].value == "1" {
			if _, err := f.query.WriteString(" IS NULL"); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(" IS NULL"); err != nil {
				return errors.Trace(err)
			}
		} else {
			if _, err := f.query.WriteString(" IS NOT NULL"); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(" IS NOT NULL"); err != nil {
				return errors.Trace(err)
			}
		}
		dontBindParam = true
		f.amountOfNotNulls++
	case filterFieldActionNotEqualTo:
		if _, err := f.query.WriteString("!="); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString("!="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionGreaterThan:
		if _, err := f.query.WriteString(">"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(">"); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionSmallerThan:
		if _, err := f.query.WriteString("<"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString("<"); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionEqualToOrGreaterThan:
		if _, err := f.query.WriteString(">="); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(">="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionEqualToOrSmallerThan:
		if _, err := f.query.WriteString("<="); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString("<="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionLike:
		if _, err := f.query.WriteString(" LIKE"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(" LIKE"); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionBetween:
		if _, err := f.query.WriteString(" BETWEEN $"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString(strconv.Itoa(int(f.incrementParameterCount()))); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}

		if _, err := f.countQuery.WriteString(" BETWEEN $"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(strconv.Itoa(int(f.parameterCount))); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}
		f.amountOfBetweens++
	}

	if !dontBindParam {
		if _, err := f.query.WriteString("?"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString("?"); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
