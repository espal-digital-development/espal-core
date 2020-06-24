package entitymutators

import (
	"strconv"

	"github.com/juju/errors"
)

func (entityMutator *EntityMutator) buildInsertQuery() error {
	fieldsLength := len(entityMutator.fields)
	if _, err := entityMutator.query.WriteString(`INSERT INTO "`); err != nil {
		return errors.Trace(err)
	}
	if _, err := entityMutator.query.WriteString(entityMutator.entity.TableName()); err != nil {
		return errors.Trace(err)
	}
	if _, err := entityMutator.query.WriteString(`" (`); err != nil {
		return errors.Trace(err)
	}
	var firstHad bool
	for k := range entityMutator.fields {
		if firstHad {
			if _, err := entityMutator.query.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstHad = true
		}
		if _, err := entityMutator.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := entityMutator.query.WriteString(entityMutator.fields[k]); err != nil {
			return errors.Trace(err)
		}
		if _, err := entityMutator.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
	}
	if _, err := entityMutator.query.WriteString(") VALUES($"); err != nil {
		return errors.Trace(err)
	}
	if _, err := entityMutator.query.WriteString(strconv.FormatUint(uint64(entityMutator.incrementParameterCount()), 10)); err != nil {
		return errors.Trace(err)
	}
	if fieldsLength > 1 {
		for i := 0; i < fieldsLength-1; i++ {
			if _, err := entityMutator.query.WriteString(",$"); err != nil {
				return errors.Trace(err)
			}
			if _, err := entityMutator.query.WriteString(strconv.FormatUint(uint64(entityMutator.incrementParameterCount()), 10)); err != nil {
				return errors.Trace(err)
			}
		}
	}
	if _, err := entityMutator.query.WriteString(")"); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (entityMutator *EntityMutator) buildUpdateQuery() error {
	if _, err := entityMutator.query.WriteString(`UPDATE "`); err != nil {
		return errors.Trace(err)
	}
	if _, err := entityMutator.query.WriteString(entityMutator.entity.TableName()); err != nil {
		return errors.Trace(err)
	}
	if _, err := entityMutator.query.WriteString(`" SET `); err != nil {
		return errors.Trace(err)
	}
	var firstHad bool
	for k := range entityMutator.fields {
		if firstHad {
			if _, err := entityMutator.query.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstHad = true
		}
		if _, err := entityMutator.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := entityMutator.query.WriteString(entityMutator.fields[k]); err != nil {
			return errors.Trace(err)
		}
		if _, err := entityMutator.query.WriteString(`"=$`); err != nil {
			return errors.Trace(err)
		}
		if _, err := entityMutator.query.WriteString(strconv.FormatUint(uint64(entityMutator.incrementParameterCount()), 10)); err != nil {
			return errors.Trace(err)
		}
	}
	if _, err := entityMutator.query.WriteString(` WHERE "id" = $`); err != nil {
		return errors.Trace(err)
	}
	if _, err := entityMutator.query.WriteString(strconv.FormatUint(uint64(entityMutator.incrementParameterCount()), 10)); err != nil {
		return errors.Trace(err)
	}

	entityMutator.values = append(entityMutator.values, entityMutator.entity.ID())

	return nil
}
