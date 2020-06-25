package entitymutators

import (
	"strconv"

	"github.com/juju/errors"
)

func (m *EntityMutator) buildInsertQuery() error {
	fieldsLength := len(m.fields)
	if _, err := m.query.WriteString(`INSERT INTO "`); err != nil {
		return errors.Trace(err)
	}
	if _, err := m.query.WriteString(m.entity.TableName()); err != nil {
		return errors.Trace(err)
	}
	if _, err := m.query.WriteString(`" (`); err != nil {
		return errors.Trace(err)
	}
	var firstHad bool
	for k := range m.fields {
		if firstHad {
			if _, err := m.query.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstHad = true
		}
		if _, err := m.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := m.query.WriteString(m.fields[k]); err != nil {
			return errors.Trace(err)
		}
		if _, err := m.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
	}
	if _, err := m.query.WriteString(") VALUES($"); err != nil {
		return errors.Trace(err)
	}
	if _, err := m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), 10)); err != nil {
		return errors.Trace(err)
	}
	if fieldsLength > 1 {
		for i := 0; i < fieldsLength-1; i++ {
			if _, err := m.query.WriteString(",$"); err != nil {
				return errors.Trace(err)
			}
			if _, err := m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), 10)); err != nil {
				return errors.Trace(err)
			}
		}
	}
	if _, err := m.query.WriteString(")"); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (m *EntityMutator) buildUpdateQuery() error {
	if _, err := m.query.WriteString(`UPDATE "`); err != nil {
		return errors.Trace(err)
	}
	if _, err := m.query.WriteString(m.entity.TableName()); err != nil {
		return errors.Trace(err)
	}
	if _, err := m.query.WriteString(`" SET `); err != nil {
		return errors.Trace(err)
	}
	var firstHad bool
	for k := range m.fields {
		if firstHad {
			if _, err := m.query.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstHad = true
		}
		if _, err := m.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := m.query.WriteString(m.fields[k]); err != nil {
			return errors.Trace(err)
		}
		if _, err := m.query.WriteString(`"=$`); err != nil {
			return errors.Trace(err)
		}
		if _, err := m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), 10)); err != nil {
			return errors.Trace(err)
		}
	}
	if _, err := m.query.WriteString(` WHERE "id" = $`); err != nil {
		return errors.Trace(err)
	}
	if _, err := m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), 10)); err != nil {
		return errors.Trace(err)
	}

	m.values = append(m.values, m.entity.ID())

	return nil
}
