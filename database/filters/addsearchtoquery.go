package filters

import (
	"strconv"

	"github.com/juju/errors"
)

func (filter *filter) addSearchToQuery() error {
	if filter.firstWhereStatementHad {
		if _, err := filter.query.WriteString(" AND ("); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" AND ("); err != nil {
			return errors.Trace(err)
		}
	} else {
		if _, err := filter.query.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
	}

	// TODO :: Longest words should come earliest?

	var firstSearchBlockHad bool
	for i := 0; i < len(filter.searchChunks); i++ {
		if firstSearchBlockHad {
			if _, err := filter.query.WriteString(" AND "); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(" AND "); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstSearchBlockHad = true
		}

		if _, err := filter.query.WriteString("("); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString("("); err != nil {
			return errors.Trace(err)
		}

		var firstSearchEntryHad bool
		for k := range filter.searchFields {
			if firstSearchEntryHad {
				if _, err := filter.query.WriteString(" OR "); err != nil {
					return errors.Trace(err)
				}
				if _, err := filter.countQuery.WriteString(" OR "); err != nil {
					return errors.Trace(err)
				}
			} else {
				firstSearchEntryHad = true
			}

			if filter.searchFields[k].TableAlias() != "" {
				if _, err := filter.query.WriteString(filter.searchFields[k].TableAlias()); err != nil {
					return errors.Trace(err)
				}
				if _, err := filter.query.WriteString("."); err != nil {
					return errors.Trace(err)
				}
			}
			if _, err := filter.query.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString(filter.searchFields[k].Name()); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString(" LIKE $"); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString(strconv.Itoa(int(filter.incrementParameterCount()))); err != nil {
				return errors.Trace(err)
			}

			if filter.searchFields[k].TableAlias() != "" {
				if _, err := filter.countQuery.WriteString(filter.searchFields[k].TableAlias()); err != nil {
					return errors.Trace(err)
				}
				if _, err := filter.countQuery.WriteString("."); err != nil {
					return errors.Trace(err)
				}
			}
			if _, err := filter.countQuery.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(filter.searchFields[k].Name()); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(" LIKE $"); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.countQuery.WriteString(strconv.Itoa(int(filter.parameterCount))); err != nil {
				return errors.Trace(err)
			}
		}

		if _, err := filter.query.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
	}

	if filter.firstWhereStatementHad {
		if _, err := filter.query.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
