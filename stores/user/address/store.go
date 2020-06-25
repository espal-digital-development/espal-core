package address

import (
	"database/sql"
	"strings"

	"github.com/espal-digital-development/espal-core/database"
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
	translationsRepository translations.Repository
	countriesRepository    countries.Repository
	loggerService          logger.Loggable
}

// ForUser fetches Addresses for userID.
func (a *AddressesStore) ForUser(userID string) (result []*Address, ok bool, err error) {
	rows, err := a.selecterDatabase.Query(`SELECT u.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "UserAddress" u
		LEFT JOIN "User" cu ON cu."id" = u."createdByID"
		LEFT JOIN "User" uu ON uu."id" = u."updatedByID"
		WHERE u."userID" = $1`, userID)
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
	result = make([]*Address, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		address := newAddress()
		err = rows.Scan(&address.id, &address.createdByID, &address.updatedByID, &address.userID, &address.createdAt, &address.updatedAt,
			&address.active, &address.firstName, &address.surname, &address.street, &address.streetLine2, &address.number,
			&address.numberAddition, &address.zipCode, &address.city, &address.state, &address.country, &address.phoneNumber, &address.email,
			&address.createdByFirstName, &address.createdBySurname, &address.updatedByFirstName, &address.updatedBySurname,
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
func (a *AddressesStore) DisplayValue(address AddressEntity, localeID uint16) string {
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
		display = a.translationsRepository.Singular(localeID, "address") + " " + address.ID()
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
		country, err := a.countriesRepository.ByID(*address.Country())
		if err != nil {
			a.loggerService.Error(errors.ErrorStack(err))
			return ""
		}
		display += " " + country.Translate(localeID)
	}

	return display
}

// GetOneByID fetches by ID.
func (a *AddressesStore) GetOneByID(id string) (*Address, bool, error) {
	result, ok, err := a.fetch(`SELECT * FROM "UserAddress" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (a *AddressesStore) GetOneByIDWithCreator(id string) (*Address, bool, error) {
	result, ok, err := a.fetch(`SELECT u.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
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
func (a *AddressesStore) Delete(ids []string) error {
	transaction, err := a.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "UserAddress" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// ToggleActive toggles the active state of the given ID(s).
func (a *AddressesStore) ToggleActive(ids []string) error {
	transaction, err := a.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Query(`UPDATE "UserAddress" SET "active" = NOT "active" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}
