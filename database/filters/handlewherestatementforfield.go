package filters

import (
	"strconv"
)

func (f *filter) handleWhereStatementForField(fieldKey int) {
	if f.firstWhereStatementHad {
		f.query.WriteString(" AND ")
		f.countQuery.WriteString(" AND ")
	} else {
		f.query.WriteString(" ")
		f.countQuery.WriteString(" ")
		f.firstWhereStatementHad = true
	}

	if f.fields[fieldKey].alias == "" {
		f.query.WriteString(f.tableAlias)
		f.query.WriteString(".")
	} else if f.fields[fieldKey].alias != "" {
		f.query.WriteString(f.fields[fieldKey].alias)
		f.query.WriteString(".")
	}

	f.query.WriteString(f.fields[fieldKey].name)
	f.countQuery.WriteString(f.fields[fieldKey].name)

	var dontBindParam bool

	switch f.fields[fieldKey].action {
	case filterFieldActionEqualTo:
		f.query.WriteString("=")
		f.countQuery.WriteString("=")
	case filterFieldActionIsNull:
		if f.fields[fieldKey].value == "1" {
			f.query.WriteString(" IS NULL")
			f.countQuery.WriteString(" IS NULL")
		} else {
			f.query.WriteString(" IS NOT NULL")
			f.countQuery.WriteString(" IS NOT NULL")
		}
		dontBindParam = true
		f.amountOfNotNulls++
	case filterFieldActionNotEqualTo:
		f.query.WriteString("!=")
		f.countQuery.WriteString("!=")
	case filterFieldActionGreaterThan:
		f.query.WriteString(">")
		f.countQuery.WriteString(">")
	case filterFieldActionSmallerThan:
		f.query.WriteString("<")
		f.countQuery.WriteString("<")
	case filterFieldActionEqualToOrGreaterThan:
		f.query.WriteString(">=")
		f.countQuery.WriteString(">=")
	case filterFieldActionEqualToOrSmallerThan:
		f.query.WriteString("<=")
		f.countQuery.WriteString("<=")
	case filterFieldActionLike:
		f.query.WriteString(" LIKE")
		f.countQuery.WriteString(" LIKE")
	case filterFieldActionBetween:
		f.query.WriteString(" BETWEEN $")
		f.query.WriteString(strconv.Itoa(int(f.incrementParameterCount())))
		f.query.WriteString(" AND ")

		f.countQuery.WriteString(" BETWEEN $")
		f.countQuery.WriteString(strconv.Itoa(int(f.parameterCount)))
		f.countQuery.WriteString(" AND ")
		f.amountOfBetweens++
	}

	if !dontBindParam {
		f.query.WriteString("?")
		f.countQuery.WriteString("?")
	}
}
