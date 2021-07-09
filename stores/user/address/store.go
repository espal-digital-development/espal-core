package address

import (
	"database/sql"
	errorsNative "errors"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/queryhelper"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

// AddressesStore data store.
type AddressesStore struct {
	selecterDatabase       database.Database
	updaterDatabase        database.Database
	deletorDatabase        database.Database
	databaseQueryHelper    queryhelper.Helper
	translationsRepository translations.Repository
	countriesRepository    countries.Repository
	loggerService          logger.Loggable
}

// ForUser fetches Addresses for userID.
// nolint:nakedret
func (s *AddressesStore) ForUser(userID string) (result []*Address, ok bool, err error) {
	rows, err := s.selecterDatabase.Query(`SELECT u.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "UserAddress" u
		LEFT JOIN "User" cu ON cu."id" = u."createdByID"
		LEFT JOIN "User" uu ON uu."id" = u."updatedByID"
		WHERE u."userID" = $1`, userID)
	if errorsNative.Is(err, sql.ErrNoRows) {
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
	result = make([]*Address, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		address := newAddress()
		err = rows.Scan(&address.id, &address.createdByID, &address.updatedByID, &address.userID, &address.createdAt,
			&address.updatedAt, &address.active, &address.firstName, &address.surname, &address.street,
			&address.streetLine2, &address.number, &address.numberAddition, &address.zipCode, &address.city,
			&address.state, &address.country, &address.phoneNumber, &address.email, &address.createdByFirstName,
			&address.createdBySurname, &address.updatedByFirstName, &address.updatedBySurname,
		)
		if err != nil {
			err = errors.Trace(err)
			return
		}
		result = append(result, address)
	}
	ok = len(result) > 0
	return
}

// DisplayValue returns a readable display representation.
func (s *AddressesStore) DisplayValue(address AddressEntity, localeID uint16) string {
	var display string
	if address.FirstName() != nil {
		display += *address.FirstName()
	}
	if address.Surname() != nil {
		if display != "" {
			display += " "
		}
		display += *address.Surname()
	}
	if display == "" {
		display = s.translationsRepository.Singular(localeID, "address") + " " + address.ID()
	}
	display += " : " + address.Street()
	if address.StreetLine2() != nil {
		display += ", " + *address.StreetLine2()
	}
	display += " " + address.Number()
	if address.NumberAddition() != nil {
		display += " " + *address.NumberAddition()
	}
	display += " " + address.City()
	if address.Country() != nil {
		country, err := s.countriesRepository.ByID(*address.Country())
		if err != nil {
			s.loggerService.Error(errors.ErrorStack(err))
			return ""
		}
		display += " " + country.Translate(localeID)
	}

	return display
}

// GetOneByID fetches by ID.
func (s *AddressesStore) GetOneByID(id string) (*Address, bool, error) {
	result, ok, err := s.fetch(`SELECT * FROM "UserAddress" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (s *AddressesStore) GetOneByIDWithCreator(id string) (*Address, bool, error) {
	result, ok, err := s.fetch(`SELECT u.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "UserAddress" u
		LEFT JOIN "User" cu ON cu."id" = u."createdByID"
		LEFT JOIN "User" uu ON uu."id" = u."updatedByID"
		WHERE u."id" = $1 LIMIT 1`, true, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// Delete deletes the given ID(s).
func (s *AddressesStore) Delete(ids []string) error {
	transaction, err := s.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	query, idsInterfaces, err := s.databaseQueryHelper.BuildDeleteWhereInIds("UserAddress", "id", ids)
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(query, idsInterfaces...); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// ToggleActive toggles the active state of the given ID(s).
func (s *AddressesStore) ToggleActive(ids []string) error {
	transaction, err := s.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	query, idsInterfaces, err := s.databaseQueryHelper.BuildUpdateWhereInIds("UserAddress",
		`SET "active" = NOT "active"`, "id", ids)
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(query, idsInterfaces...); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}
