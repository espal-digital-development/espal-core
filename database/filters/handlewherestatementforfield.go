package filters

import (
	"strconv"

	"github.com/juju/errors"
)

func (filter *filter) handleWhereStatementForField(fieldKey int) error {
	if filter.firstWhereStatementHad {
		if _, err := filter.query.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}
	} else {
		if _, err := filter.query.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		filter.firstWhereStatementHad = true
	}

	if filter.fields[fieldKey].alias == "" {
		if _, err := filter.query.WriteString(filter.tableAlias); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString("."); err != nil {
			return errors.Trace(err)
		}
	} else if filter.fields[fieldKey].alias != "" {
		if _, err := filter.query.WriteString(filter.fields[fieldKey].alias); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString("."); err != nil {
			return errors.Trace(err)
		}
	}

	if _, err := filter.query.WriteString(filter.fields[fieldKey].name); err != nil {
		return errors.Trace(err)
	}
	if _, err := filter.countQuery.WriteString(filter.fields[fieldKey].name); err != nil {
		return errors.Trace(err)
	}

	var dontBindParam bool

	switch filter.fields[fieldKey].action {
	case filterFieldActionEqualTo:
		if _, err := filter.query.WriteString("="); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString("="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionIsNull:
		if filter.fields[fieldKey].value == "1" {
			if _, err := filter.query.WriteString(" IS NULL"); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(" IS NULL"); err != nil {
				return errors.Trace(err)
			}
		} else {
			if _, err := filter.query.WriteString(" IS NOT NULL"); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(" IS NOT NULL"); err != nil {
				return errors.Trace(err)
			}
		}
		dontBindParam = true
		filter.amountOfNotNulls++
	case filterFieldActionNotEqualTo:
		if _, err := filter.query.WriteString("!="); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString("!="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionGreaterThan:
		if _, err := filter.query.WriteString(">"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(">"); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionSmallerThan:
		if _, err := filter.query.WriteString("<"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString("<"); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionEqualToOrGreaterThan:
		if _, err := filter.query.WriteString(">="); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(">="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionEqualToOrSmallerThan:
		if _, err := filter.query.WriteString("<="); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString("<="); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionLike:
		if _, err := filter.query.WriteString(" LIKE"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" LIKE"); err != nil {
			return errors.Trace(err)
		}
	case filterFieldActionBetween:
		if _, err := filter.query.WriteString(" BETWEEN $"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(strconv.Itoa(int(filter.incrementParameterCount()))); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}

		if _, err := filter.countQuery.WriteString(" BETWEEN $"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(strconv.Itoa(int(filter.parameterCount))); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" AND "); err != nil {
			return errors.Trace(err)
		}
		filter.amountOfBetweens++
	}

	if !dontBindParam {
		if _, err := filter.query.WriteString("?"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString("?"); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
