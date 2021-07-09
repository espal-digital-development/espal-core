package entitymutators

import (
	"strconv"

	"github.com/espal-digital-development/system/units"
)

func (m *EntityMutator) buildInsertQuery() {
	fieldsLength := len(m.fields)
	m.query.WriteString(`INSERT INTO "`)
	m.query.WriteString(m.entity.TableName())
	m.query.WriteString(`" (`)
	var firstHad bool
	for k := range m.fields {
		if firstHad {
			m.query.WriteString(",")
		} else {
			firstHad = true
		}
		m.query.WriteString(`"`)
		m.query.WriteString(m.fields[k])
		m.query.WriteString(`"`)
	}
	m.query.WriteString(") VALUES($")
	m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), units.Base10))
	if fieldsLength > 1 {
		for i := 0; i < fieldsLength-1; i++ {
			m.query.WriteString(",$")
			m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), units.Base10))
		}
	}
	m.query.WriteString(")")
}

func (m *EntityMutator) buildUpdateQuery() {
	m.query.WriteString(`UPDATE "`)
	m.query.WriteString(m.entity.TableName())
	m.query.WriteString(`" SET `)
	var firstHad bool
	for k := range m.fields {
		if firstHad {
			m.query.WriteString(",")
		} else {
			firstHad = true
		}
		m.query.WriteString(`"`)
		m.query.WriteString(m.fields[k])
		m.query.WriteString(`"=$`)
		m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), units.Base10))
	}
	m.query.WriteString(` WHERE "id" = $`)
	m.query.WriteString(strconv.FormatUint(uint64(m.incrementParameterCount()), units.Base10))

	m.values = append(m.values, m.entity.ID())
}
