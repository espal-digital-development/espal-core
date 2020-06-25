package contact

import (
	"database/sql"
	"strings"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

// ContactsStore data store.
type ContactsStore struct {
	selecterDatabase       database.Database
	deletorDatabase        database.Database
	translationsRepository translations.Repository
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (c *ContactsStore) GetOneByIDWithCreator(id string) (*Contact, bool, error) {
	result, ok, err := c.fetch(`SELECT u.*, c."firstName", c."surname", cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "UserContact" u
		LEFT JOIN "User" c ON c."id" = u."contactID"
		LEFT JOIN "User" cu ON cu."id" = u."createdByID"
		LEFT JOIN "User" uu ON uu."id" = u."updatedByID"
		WHERE u."id" = $1 LIMIT 1`, true, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// Delete deletes the given ID(s).
func (c *ContactsStore) Delete(ids []string) error {
	transaction, err := c.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "UserContact" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// Name returns the presentable name for the User's Contact.
func (c *ContactsStore) Name(contact *Contact, languageID uint16) string {
	var name string
	if contact.ContactFirstName() != nil {
		name = *contact.ContactFirstName()
	}
	if contact.ContactSurname() != nil {
		name += " " + *contact.ContactSurname()
	}
	if name == "" {
		name = c.translationsRepository.Singular(languageID, "user") + " " + contact.ID()
	}
	return name
}

// ForUser fetches UserContacts for userID.
func (c *ContactsStore) ForUser(userID string) (result []*Contact, ok bool, err error) {
	rows, err := c.selecterDatabase.Query(`SELECT u.*, c."firstName", c."surname", cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "UserContact" u
		LEFT JOIN "User" c ON c."id" = u."contactID"
		LEFT JOIN "User" cu ON cu."id" = u."createdByID"
		LEFT JOIN "User" uu ON uu."id" = u."updatedByID"
		WHERE u."userID" = $1
		ORDER BY u."sorting"`, userID)
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
	result = make([]*Contact, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		contact := newContact()
		err = rows.Scan(&contact.id, &contact.createdByID, &contact.updatedByID, &contact.userID, &contact.contactID, &contact.createdAt,
			&contact.updatedAt, &contact.sorting, &contact.comments, &contact.contactFirstName, &contact.contactSurname,
			&contact.createdByFirstName, &contact.createdBySurname, &contact.updatedByFirstName, &contact.updatedBySurname,
		)
		if err != nil {
			err = errors.Trace(err)
			return
		}
		result = append(result, contact)
	}
	ok = len(result) > 0
	return
}
