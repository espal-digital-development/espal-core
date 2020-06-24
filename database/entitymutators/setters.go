package entitymutators

import (
	"time"

	"github.com/juju/errors"
)

// SetBool checks and adds bool value based on the given field.
func (entityMutator *EntityMutator) SetBool(field string, value bool, originalValue bool) {
	if entityMutator.entity.ID() == "" || value != originalValue {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	}
}

// SetUint checks and adds uint value based on the given field.
func (entityMutator *EntityMutator) SetUint(field string, value uint, originalValue uint) {
	if entityMutator.entity.ID() == "" || value != originalValue {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	}
}

// SetNullableUint checks and adds *uint value based on the given field.
func (entityMutator *EntityMutator) SetNullableUint(field string, value *uint, originalValue *uint) {
	if entityMutator.entity.ID() == "" && value != nil && *value > 0 {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	} else if entityMutator.entity.ID() != "" {
		if value != nil && *value > 0 && (originalValue == nil || *value != *originalValue) {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, value)
		} else if (value == nil || *value == 0) && originalValue != nil {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, nil)
		}
	}
}

// SetUint16 checks and adds uint16 value based on the given field.
func (entityMutator *EntityMutator) SetUint16(field string, value uint16, originalValue uint16) {
	if entityMutator.entity.ID() == "" || value != originalValue {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	}
}

// SetNullableUint16 checks and adds *uint16 value based on the given field.
func (entityMutator *EntityMutator) SetNullableUint16(field string, value *uint16, originalValue *uint16) {
	if entityMutator.entity.ID() == "" && value != nil && *value > 0 {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	} else if entityMutator.entity.ID() != "" {
		if value != nil && *value > 0 && (originalValue == nil || *value != *originalValue) {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, value)
		} else if (value == nil || *value == 0) && originalValue != nil {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, nil)
		}
	}
}

// SetString checks and adds string value based on the given field.
func (entityMutator *EntityMutator) SetString(field string, value *string, originalValue string) {
	if entityMutator.entity.ID() == "" || *value != originalValue {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, *value)
	}
}

// SetNullableString checks and adds *string value based on the given field.
func (entityMutator *EntityMutator) SetNullableString(field string, value *string, originalValue *string) {
	if entityMutator.entity.ID() == "" && value != nil && *value != "" {
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	} else if entityMutator.entity.ID() != "" {
		if value != nil && *value != "" && (originalValue == nil || *value != *originalValue) {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, value)
		} else if (value != nil || *value == "") && originalValue != nil {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, nil)
		}
	}
}

// SetNullableTime checks and adds *string value based on the given field.
// A string is given, to more easily parse and process the logic from forms.
func (entityMutator *EntityMutator) SetNullableTime(field string, value *string, originalValue *time.Time) error {
	if entityMutator.entity.ID() == "" && *value != "" {
		_, err := time.Parse(time.RFC3339[0:9], *value)
		if err != nil {
			return errors.Trace(err)
		}
		entityMutator.fields = append(entityMutator.fields, field)
		entityMutator.values = append(entityMutator.values, value)
	} else if entityMutator.entity.ID() != "" {
		if *value != "" && originalValue == nil {
			_, err := time.Parse(time.RFC3339[0:9],
				*value)
			if err != nil {
				return errors.Trace(err)
			}
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, value)
		} else if *value == "" && originalValue != nil {
			entityMutator.fields = append(entityMutator.fields, field)
			entityMutator.values = append(entityMutator.values, nil)
		}
	}
	return nil
}
