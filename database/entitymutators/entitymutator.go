package entitymutators

import (
	"bytes"
	"strings"
	"time"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/juju/errors"
)

var _ Mutator = &EntityMutator{}

type entity interface {
	ID() string
	TableName() string
}

type form interface {
	FieldValue(string) string
}

// Mutator represents an object that runs a database mutation.
type Mutator interface {
	Execute(context contexts.Context) error
	SetCustomReturnPath(returnPath string)
	SetExtraURLQueryParams(paramsString string)
	GetInsertedOrUpdatedID() string
	RedirectURL() string
	SetBool(field string, value bool, originalValue bool)
	SetUint(field string, value uint, originalValue uint)
	SetNullableUint(field string, value *uint, originalValue *uint)
	SetUint16(field string, value uint16, originalValue uint16)
	SetNullableUint16(field string, value *uint16, originalValue *uint16)
	SetString(field string, value *string, originalValue string)
	SetNullableString(field string, value *string, originalValue *string)
	SetNullableTime(field string, value *string, originalValue *time.Time) error
}

// EntityMutator holds a single database mutation run.
type EntityMutator struct {
	inserterDatabase database.Database
	updaterDatabase  database.Database
	entity           entity
	lastInsertedID   string
	path             string
	returnPath       string
	extraQueryParams string
	formAction       string
	query            *bytes.Buffer
	fields           []string
	values           []interface{}

	// parameterCount counts all the incremental $1, $2 etc. parameters that are used
	parameterCount uint16
}

// Execute is a shortcut to quickly run insert/update actions
// for Admin Create/Update routes.
func (entityMutator *EntityMutator) Execute(context contexts.Context) error {
	var err error
	fieldsLength := len(entityMutator.fields)
	if fieldsLength == 0 {
		return context.SetFlashInfoMessage(context.Translate("noDataWasChanged"))
	}

	user, ok, err := context.GetUser()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		return errors.Errorf("user couldn't be retrieved from the context")
	}

	if entityMutator.entity.ID() == "" {
		entityMutator.fields = append(entityMutator.fields, "createdByID")
	} else {
		entityMutator.fields = append(entityMutator.fields, "updatedByID")
	}
	entityMutator.values = append(entityMutator.values, user.ID())

	entityMutator.query = new(bytes.Buffer)
	if entityMutator.entity.ID() == "" {
		if err := entityMutator.buildInsertQuery(); err != nil {
			return errors.Trace(err)
		}
	} else {
		if err := entityMutator.buildUpdateQuery(); err != nil {
			return errors.Trace(err)
		}
	}

	var row database.Row
	var lastInsertedID string

	if entityMutator.entity.ID() == "" {
		row = entityMutator.inserterDatabase.QueryRow(entityMutator.query.String()+` RETURNING "id"`, entityMutator.values...)
	} else {
		row = entityMutator.updaterDatabase.QueryRow(entityMutator.query.String()+` RETURNING "id"`, entityMutator.values...)
	}
	if err := row.Scan(&lastInsertedID); err != nil {
		return errors.Trace(err)
	}

	if err != nil {
		if entityMutator.entity.ID() == "" {
			if err := context.SetFlashErrorMessage(context.Translate("creationHasFailed") + ": " + err.Error()); err != nil {
				return errors.Trace(err)
			}
		} else {
			if err := context.SetFlashErrorMessage(context.Translate("updateHasFailed") + ": " + err.Error()); err != nil {
				return errors.Trace(err)
			}
		}
		return errors.Trace(err)
	}

	if entityMutator.entity.ID() == "" && lastInsertedID == "" {
		return errors.Errorf("lastInsertedID was not set")
	}

	entityMutator.lastInsertedID = lastInsertedID

	if entityMutator.entity.ID() == "" {
		if err := context.SetFlashSuccessMessage(context.Translate("creationWasSuccessful")); err != nil {
			return errors.Trace(err)
		}
	} else {
		if err := context.SetFlashSuccessMessage(context.Translate("updateWasSuccessful")); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

// incrementParameterCount increments the parameter count and then returns the new value.
func (entityMutator *EntityMutator) incrementParameterCount() uint16 {
	entityMutator.parameterCount++
	return entityMutator.parameterCount
}

// SetCustomReturnPath sets a custom return path to redirect to
func (entityMutator *EntityMutator) SetCustomReturnPath(returnPath string) {
	entityMutator.returnPath = returnPath
}

// SetExtraURLQueryParams adds extra query parameters to the RedirectURL call.
// No need to add prefixed `?` or `&`
func (entityMutator *EntityMutator) SetExtraURLQueryParams(paramsString string) {
	entityMutator.extraQueryParams = strings.TrimLeft(paramsString, "&")
}

// GetInsertedOrUpdatedID returns the last inserted ID.
func (entityMutator *EntityMutator) GetInsertedOrUpdatedID() string {
	if entityMutator.entity.ID() != "" {
		return entityMutator.entity.ID()
	}
	return entityMutator.lastInsertedID
}

// RedirectURL returns where the
func (entityMutator *EntityMutator) RedirectURL() string {
	var url string
	var skipParams bool

	switch {
	case validators.SaveAndCreate == entityMutator.formAction:
		url = "/" + entityMutator.path + "/Create"
	case validators.SaveAndClone == entityMutator.formAction:
		if entityMutator.entity.ID() != "" {
			// TODO :: Clone; forward post-data without actually posting?
			//         This doesn't work yet. Also fields like FileField give issues.
			url = "/Create"
		}
	case entityMutator.returnPath != "":
		url = "/" + entityMutator.returnPath
		skipParams = true
	default:
		url = "/" + entityMutator.path
	}

	if !skipParams && entityMutator.extraQueryParams != "" {
		if !strings.Contains(url, "?") {
			url += "?"
		} else {
			url += "&"
		}
		url += entityMutator.extraQueryParams
	}

	return url
}
