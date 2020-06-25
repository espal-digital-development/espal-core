package entitymutators

import (
	"time"

	"github.com/juju/errors"
)

// SetBool checks and adds bool value based on the given field.
func (m *EntityMutator) SetBool(field string, value bool, originalValue bool) {
	if m.entity.ID() == "" || value != originalValue {
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	}
}

// SetUint checks and adds uint value based on the given field.
func (m *EntityMutator) SetUint(field string, value uint, originalValue uint) {
	if m.entity.ID() == "" || value != originalValue {
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	}
}

// SetNullableUint checks and adds *uint value based on the given field.
func (m *EntityMutator) SetNullableUint(field string, value *uint, originalValue *uint) {
	if m.entity.ID() == "" && value != nil && *value > 0 {
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	} else if m.entity.ID() != "" {
		if value != nil && *value > 0 && (originalValue == nil || *value != *originalValue) {
			m.fields = append(m.fields, field)
			m.values = append(m.values, value)
		} else if (value == nil || *value == 0) && originalValue != nil {
			m.fields = append(m.fields, field)
			m.values = append(m.values, nil)
		}
	}
}

// SetUint16 checks and adds uint16 value based on the given field.
func (m *EntityMutator) SetUint16(field string, value uint16, originalValue uint16) {
	if m.entity.ID() == "" || value != originalValue {
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	}
}

// SetNullableUint16 checks and adds *uint16 value based on the given field.
func (m *EntityMutator) SetNullableUint16(field string, value *uint16, originalValue *uint16) {
	if m.entity.ID() == "" && value != nil && *value > 0 {
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	} else if m.entity.ID() != "" {
		if value != nil && *value > 0 && (originalValue == nil || *value != *originalValue) {
			m.fields = append(m.fields, field)
			m.values = append(m.values, value)
		} else if (value == nil || *value == 0) && originalValue != nil {
			m.fields = append(m.fields, field)
			m.values = append(m.values, nil)
		}
	}
}

// SetString checks and adds string value based on the given field.
func (m *EntityMutator) SetString(field string, value *string, originalValue string) {
	if m.entity.ID() == "" || *value != originalValue {
		m.fields = append(m.fields, field)
		m.values = append(m.values, *value)
	}
}

// SetNullableString checks and adds *string value based on the given field.
func (m *EntityMutator) SetNullableString(field string, value *string, originalValue *string) {
	if m.entity.ID() == "" && value != nil && *value != "" {
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	} else if m.entity.ID() != "" {
		if value != nil && *value != "" && (originalValue == nil || *value != *originalValue) {
			m.fields = append(m.fields, field)
			m.values = append(m.values, value)
		} else if (value != nil || *value == "") && originalValue != nil {
			m.fields = append(m.fields, field)
			m.values = append(m.values, nil)
		}
	}
}

// SetNullableTime checks and adds *string value based on the given field.
// A string is given, to more easily parse and process the logic from forms.
func (m *EntityMutator) SetNullableTime(field string, value *string, originalValue *time.Time) error {
	if m.entity.ID() == "" && *value != "" {
		_, err := time.Parse(time.RFC3339[0:9], *value)
		if err != nil {
			return errors.Trace(err)
		}
		m.fields = append(m.fields, field)
		m.values = append(m.values, value)
	} else if m.entity.ID() != "" {
		if *value != "" && originalValue == nil {
			_, err := time.Parse(time.RFC3339[0:9],
				*value)
			if err != nil {
				return errors.Trace(err)
			}
			m.fields = append(m.fields, field)
			m.values = append(m.values, value)
		} else if *value == "" && originalValue != nil {
			m.fields = append(m.fields, field)
			m.values = append(m.values, nil)
		}
	}
	return nil
}
