package filters

import (
	"strconv"

	"github.com/juju/errors"
)

// nolint:gocyclo,funlen
func (f *filter) addSearchToQuery() error {
	if f.firstWhereStatementHad {
		if _, err := f.query.WriteString(" AND ("); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(" AND ("); err != nil {
			return errors.Trace(err)
		}
	} else {
		if _, err := f.query.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
	}

	// TODO :: Longest words should come earliest?

	var firstSearchBlockHad bool
	for i := 0; i < len(f.searchChunks); i++ {
		if firstSearchBlockHad {
			if _, err := f.query.WriteString(" AND "); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(" AND "); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstSearchBlockHad = true
		}

		if _, err := f.query.WriteString("("); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString("("); err != nil {
			return errors.Trace(err)
		}

		var firstSearchEntryHad bool
		for k := range f.searchFields {
			if firstSearchEntryHad {
				if _, err := f.query.WriteString(" OR "); err != nil {
					return errors.Trace(err)
				}
				if _, err := f.countQuery.WriteString(" OR "); err != nil {
					return errors.Trace(err)
				}
			} else {
				firstSearchEntryHad = true
			}

			if f.searchFields[k].TableAlias() != "" {
				if _, err := f.query.WriteString(f.searchFields[k].TableAlias()); err != nil {
					return errors.Trace(err)
				}
				if _, err := f.query.WriteString("."); err != nil {
					return errors.Trace(err)
				}
			}
			if _, err := f.query.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.query.WriteString(f.searchFields[k].Name()); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.query.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.query.WriteString(" LIKE $"); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.query.WriteString(strconv.Itoa(int(f.incrementParameterCount()))); err != nil {
				return errors.Trace(err)
			}

			if f.searchFields[k].TableAlias() != "" {
				if _, err := f.countQuery.WriteString(f.searchFields[k].TableAlias()); err != nil {
					return errors.Trace(err)
				}
				if _, err := f.countQuery.WriteString("."); err != nil {
					return errors.Trace(err)
				}
			}
			if _, err := f.countQuery.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(f.searchFields[k].Name()); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(`"`); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(" LIKE $"); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.countQuery.WriteString(strconv.Itoa(int(f.parameterCount))); err != nil {
				return errors.Trace(err)
			}
		}

		if _, err := f.query.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
	}

	if f.firstWhereStatementHad {
		if _, err := f.query.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.countQuery.WriteString(")"); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
