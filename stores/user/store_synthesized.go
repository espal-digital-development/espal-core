// Code generated by espal-store-synthesizer. DO NOT EDIT.
package user

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/juju/errors"
)

var _ Store = &UsersStore{}

// Store represents a data interaction object.
type Store interface {
	GetOne(id string) (*User, bool, error)
	GetOneActive(id string) (*User, bool, error)
	GetOneByIDWithCreator(id string) (*User, bool, error)
	GetOneByEmail(email string) (*User, bool, error)
	GetOneActiveByEmail(email string) (*User, bool, error)
	GetOneIDAndPasswordForActiveByEmail(email string) (*User, bool, error)
	GetOneIDForActivationHash(hash string) (string, bool, error)
	ExistsByEmail(email string) (bool, error)
	SetPasswordResetHashForUser(id string, hash string) error
	SetPasswordForUser(id string, password []byte) error
	Activate(id string) error
	GetAvatar(id string) (*string, bool, error)
	UnsetAvatar(id string) error
	Delete(ids []string) error
	ToggleActive(ids []string) error
	Register(email string, password []byte, firstName *string, surname *string,
		languageID uint16) (string, error)
	RecoverWithNewPassword(id string, password []byte, resetCount uint8) error
	HasUserRight(userEntity UserEntity, userRightName string) (bool, error)
	Name(user UserEntity, languageID uint16) string
	NameWithEmail(user UserEntity, languageID uint16) string
	Search(context filters.QueryReader) (result []*User, filter filters.Filter, err error)
	Filter(context filters.QueryReader) (result []*User, filter filters.Filter, err error)
}

func (u *UsersStore) fetch(query string, withCreators bool, params ...interface{}) (result []*User, ok bool, err error) {
	rows, err := u.selecterDatabase.Query(query, params...)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	defer func(dbRows database.Rows) {
		closeErr := dbRows.Close()
		if err != nil && closeErr != nil {
			err = errors.Wrap(err, closeErr)
		} else if closeErr != nil {
			err = errors.Trace(closeErr)
		}
	}(rows)
	result = make([]*User, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		u := newUser()
		fields := []interface{}{&u.id, &u.createdByID, &u.updatedByID, &u.defaultDeliveryAddressID, &u.defaultInvoiceAddressID, &u.createdAt, &u.updatedAt, &u.active, &u.country, &u.language, &u.firstName, &u.surname, &u.dateOfBirth, &u.email, &u.password, &u.avatar, &u.priority, &u.activationHash, &u.activatedAt, &u.passwordResetHash, &u.passwordResetLastSendAt, &u.passwordLastResetAt, &u.passwordResetCount, &u.biography, &u.comments, &u.currencies}
		if withCreators {
			fields = append(fields, &u.createdByFirstName, &u.createdBySurname, &u.updatedByFirstName, &u.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, u)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of UsersStore.
func New(selecterDatabase database.Database, inserterDatabase database.Database, updaterDatabase database.Database, deletorDatabase database.Database, databaseFiltersFactory filters.Factory, translationsRepository translations.Repository, userRightsRepository userrights.Repository) (*UsersStore, error) {
	u := &UsersStore{
		selecterDatabase:       selecterDatabase,
		inserterDatabase:       inserterDatabase,
		updaterDatabase:        updaterDatabase,
		deletorDatabase:        deletorDatabase,
		databaseFiltersFactory: databaseFiltersFactory,
		translationsRepository: translationsRepository,
		userRightsRepository:   userRightsRepository,
	}
	if err := u.buildQueries(); err != nil {
		return nil, errors.Trace(err)
	}
	return u, nil
}
