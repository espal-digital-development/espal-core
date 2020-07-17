package filters

import (
	"strconv"
)

func (f *filter) addSearchToQuery() {
	if f.firstWhereStatementHad {
		f.query.WriteString(" AND (")
		f.countQuery.WriteString(" AND (")
	} else {
		f.query.WriteString(" ")
		f.countQuery.WriteString(" ")
	}

	// TODO :: Longest words should come earliest?

	var firstSearchBlockHad bool
	for i := 0; i < len(f.searchChunks); i++ {
		if firstSearchBlockHad {
			f.query.WriteString(" AND ")
			f.countQuery.WriteString(" AND ")
		} else {
			firstSearchBlockHad = true
		}

		f.query.WriteString("(")
		f.countQuery.WriteString("(")

		var firstSearchEntryHad bool
		for k := range f.searchFields {
			if firstSearchEntryHad {
				f.query.WriteString(" OR ")
				f.countQuery.WriteString(" OR ")
			} else {
				firstSearchEntryHad = true
			}

			if f.searchFields[k].TableAlias() != "" {
				f.query.WriteString(f.searchFields[k].TableAlias())
				f.query.WriteString(".")
			}
			f.query.WriteString(`"`)
			f.query.WriteString(f.searchFields[k].Name())
			f.query.WriteString(`"`)
			f.query.WriteString(" LIKE $")
			f.query.WriteString(strconv.Itoa(int(f.incrementParameterCount())))

			if f.searchFields[k].TableAlias() != "" {
				f.countQuery.WriteString(f.searchFields[k].TableAlias())
				f.countQuery.WriteString(".")
			}
			f.countQuery.WriteString(`"`)
			f.countQuery.WriteString(f.searchFields[k].Name())
			f.countQuery.WriteString(`"`)
			f.countQuery.WriteString(" LIKE $")
			f.countQuery.WriteString(strconv.Itoa(int(f.parameterCount)))
		}

		f.query.WriteString(")")
		f.countQuery.WriteString(")")
	}

	if f.firstWhereStatementHad {
		f.query.WriteString(")")
		f.countQuery.WriteString(")")
	}
}
