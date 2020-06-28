package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/espal-digital-development/espal-core/text"
	"github.com/juju/errors"
)

const (
	// stubUUID is used for the cyclical createdByID field, which can only
	// be added later when the User registers him-/herself.
	stubUUID   = "ffffffff-ffff-ffff-ffff-ffffffffffff"
	hashLength = 72
)

var errBuildQueriesMultipleTimes = errors.New("cannot run buildQueries multiple times")

// UsersStore data store.
type UsersStore struct {
	selecterDatabase       database.Database
	inserterDatabase       database.Database
	updaterDatabase        database.Database
	deletorDatabase        database.Database
	databaseFiltersFactory filters.Factory
	translationsRepository translations.Repository
	userRightsRepository   userrights.Repository

	queries map[string]string
}

func (u *UsersStore) buildQueries() error {
	if u.queries != nil {
		return errors.Trace(errBuildQueriesMultipleTimes)
	}
	name := (&User{}).TableName()
	alias := (&User{}).TableAlias()
	u.queries = map[string]string{
		"GetOne": fmt.Sprintf(`SELECT * FROM "%s" WHERE "id" = $1 LIMIT 1`, name),
		"GetOneActive": fmt.Sprintf(`SELECT * FROM "%s" WHERE "id" = $1 `+
			`AND "active" = true LIMIT 1`, name),
		"GetOneByIDWithCreator": fmt.Sprintf(`SELECT %s.*, cu."firstName", cu."surname", `+
			`uu."firstName", uu."surname" FROM "%s" %s LEFT JOIN "User" cu ON cu."id" = %s."createdByID" `+
			`LEFT JOIN "User" uu ON uu."id" = %s."updatedByID" WHERE %s."id" = $1 LIMIT 1`,
			alias, name, alias, alias, alias, alias),
		"GetOneByEmail": fmt.Sprintf(`SELECT * FROM "%s" WHERE "email" = $1 LIMIT 1`, name),
		"GetOneActiveByEmail": fmt.Sprintf(`SELECT * FROM "%s" WHERE "email" = $1 `+
			`AND "active" = true LIMIT 1`, name),
		"GetOneIDAndPasswordForActiveByEmail": fmt.Sprintf(`SELECT "id", "password" FROM "%s" `+
			`WHERE "email" = $1 AND "active" = true LIMIT 1`, name),
		"GetOneIDForActivationHash":   fmt.Sprintf(`SELECT "id" FROM "%s" WHERE "activationHash" = $1 LIMIT 1`, name),
		"ExistsByEmail":               fmt.Sprintf(`SELECT 1 FROM "%s" WHERE "email" = $1 LIMIT 1`, name),
		"SetPasswordResetHashForUser": fmt.Sprintf(`UPDATE "%s" SET "passwordResetHash" = $1 WHERE "id" = $2`, name),
		"SetPasswordForUser":          fmt.Sprintf(`UPDATE "%s" SET "password" = $1 WHERE "id" = $2`, name),
		"Activate": fmt.Sprintf(`UPDATE "%s" SET "activatedAt" = NOW(), "activationhash" = NULL, `+
			`"active" = true WHERE "id" = $1`, name),
		"GetAvatar":           fmt.Sprintf(`SELECT "avatar" FROM "%s" WHERE "id" = $1 LIMIT 1`, name),
		"UnsetAvatar":         fmt.Sprintf(`UPDATE "%s" SET "avatar" = NULL WHERE "id" = $1`, name),
		"DeleteUserGroupUser": fmt.Sprintf(`DELETE FROM "UserGroupUser" WHERE "userID" IN ('%%s')`),
		"Delete":              fmt.Sprintf(`DELETE FROM "%s" WHERE "id" IN ('%%s')`, name),
		"ToggleActive":        fmt.Sprintf(`UPDATE "%s" SET "active" = NOT "active" WHERE "id" IN ('%%s')`, name),
		"Register": fmt.Sprintf(`INSERT INTO "%s"("createdByID","language","email","password",`+
			`"activationHash","firstName","surname") VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`, name),
		"RegisterUpdate": `UPDATE "User" SET "createdByID" = $1 WHERE "id" = $2 LIMIT 1`,
		"RecoverWithNewPassword": fmt.Sprintf(`UPDATE "%s" SET "passwordResetHash" = NULL, "password" = $1, `+
			`"passwordLastResetAt" = NOW(), "passwordResetCount" = $2 WHERE "id" = $3`, name),
		"HasUserRight": `SELECT 1 FROM "UserGroupUser" uu
		JOIN "UserGroup" ug ON ug."id" = uu."userGroupID" AND $1:::STRING = ANY (string_to_array(ug."userRights",','))
		WHERE uu."userID" = $2 LIMIT 1`,
	}
	return nil
}

// GetOne fetches by ID.
func (u *UsersStore) GetOne(id string) (*User, bool, error) {
	result, ok, err := u.fetch(u.queries["GetOne"], false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneActive fetches by ID and must be active.
func (u *UsersStore) GetOneActive(id string) (*User, bool, error) {
	result, ok, err := u.fetch(u.queries["GetOneActive"], false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (u *UsersStore) GetOneByIDWithCreator(id string) (*User, bool, error) {
	result, ok, err := u.fetch(u.queries["GetOneByIDWithCreator"], true, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneByEmail fetches by Email.
func (u *UsersStore) GetOneByEmail(email string) (*User, bool, error) {
	result, ok, err := u.fetch(u.queries["GetOneByEmail"], false, email)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneActiveByEmail fetches by Email and must be Active.
func (u *UsersStore) GetOneActiveByEmail(email string) (*User, bool, error) {
	result, ok, err := u.fetch(u.queries["GetOneActiveByEmail"], false, email)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneIDAndPasswordForActiveByEmail only fetches the ID and Password parts into the User result
// based on the requested Email. The User must be marked Active in the database.
func (u *UsersStore) GetOneIDAndPasswordForActiveByEmail(email string) (*User, bool, error) {
	user := newUser()
	err := u.selecterDatabase.QueryRow(u.queries["GetOneIDAndPasswordForActiveByEmail"], email).
		Scan(&user.id, &user.password)
	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, errors.Trace(err)
	}
	return user, true, nil
}

// GetOneIDForActivationHash only fetches the ID parts into the User result
// basde on the requetsed hash.
func (u *UsersStore) GetOneIDForActivationHash(hash string) (string, bool, error) {
	var id string
	err := u.selecterDatabase.QueryRow(u.queries["GetOneIDForActivationHash"], hash).Scan(&id)
	if err == sql.ErrNoRows {
		return "", false, nil
	}
	if err != nil {
		return "", false, errors.Trace(err)
	}
	return id, true, nil
}

// ExistsByEmail will check and determine if the requested User with the given
// email address exists.
func (u *UsersStore) ExistsByEmail(email string) (bool, error) {
	var exists bool
	err := u.selecterDatabase.QueryRow(u.queries["ExistsByEmail"], email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Trace(err)
	}
	if err == sql.ErrNoRows || !exists {
		return false, nil
	}
	return true, nil
}

// SetPasswordResetHashForUser saves the PasswordResetHash for the given User ID.
func (u *UsersStore) SetPasswordResetHashForUser(id string, hash string) error {
	if _, err := u.updaterDatabase.Exec(u.queries["SetPasswordResetHashForUser"], hash, id); err != nil {
		return errors.Trace(err)
	}
	return nil
}

// SetPasswordForUser sets the password for the given User ID.
func (u *UsersStore) SetPasswordForUser(id string, password []byte) error {
	_, err := u.updaterDatabase.Query(u.queries["SetPasswordForUser"], password, id)
	return errors.Trace(err)
}

// Activate will activate the given User id and clear the activation hash.
func (u *UsersStore) Activate(id string) error {
	_, err := u.updaterDatabase.Exec(u.queries["Activate"], id)
	return errors.Trace(err)
}

// GetAvatar returns the avatar for the given User ID.
func (u *UsersStore) GetAvatar(id string) (*string, bool, error) {
	var avatar *string
	err := u.selecterDatabase.QueryRow(u.queries["GetAvatar"], id).Scan(&avatar)
	if err != nil && err != sql.ErrNoRows {
		return nil, false, errors.Trace(err)
	}
	if err == sql.ErrNoRows || avatar == nil {
		return nil, false, nil
	}
	return avatar, true, nil
}

// UnsetAvatar will unset the User's avatar.
func (u *UsersStore) UnsetAvatar(id string) error {
	_, err := u.updaterDatabase.Exec(u.queries["UnsetAvatar"], id)
	return errors.Trace(err)
}

// Delete deletes the given ID(s).
func (u *UsersStore) Delete(ids []string) error {
	transaction, err := u.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(fmt.Sprintf(u.queries["DeleteUserGroupUser"], strings.Join(ids, "','"))); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(fmt.Sprintf(u.queries["Delete"], strings.Join(ids, "','"))); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// ToggleActive toggles the active state of the given ID(s).
func (u *UsersStore) ToggleActive(ids []string) error {
	transaction, err := u.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Query(fmt.Sprintf(u.queries["ToggleActive"], strings.Join(ids, "','"))); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// Register registers a new User with the given base information.
func (u *UsersStore) Register(email string, password []byte, firstName *string, surname *string,
	languageID uint16) (string, error) {
	activationHash := text.RandomString(hashLength)
	var insertedID string
	row := u.inserterDatabase.QueryRow(u.queries["Register"],
		stubUUID, languageID, email, password, activationHash, firstName, surname)
	if err := row.Scan(&insertedID); err != nil {
		return "", errors.Trace(err)
	}
	_, err := u.updaterDatabase.Exec(u.queries["RegisterUpdate"], insertedID, insertedID)
	return activationHash, errors.Trace(err)
}

// RecoverWithNewPassword will set the password as recovered and increments the resetCount.
func (u *UsersStore) RecoverWithNewPassword(id string, password []byte, resetCount uint8) error {
	_, err := u.updaterDatabase.Exec(u.queries["RecoverWithNewPassword"], password, resetCount, id)
	return errors.Trace(err)
}

// HasUserRight checks if the userright is present for this User.
func (u *UsersStore) HasUserRight(userEntity UserEntity, userRightName string) (bool, error) {
	userRight, err := u.userRightsRepository.GetCode(userRightName)
	if err != nil {
		return false, errors.Trace(err)
	}

	var hasUserRight uint8
	err = u.selecterDatabase.QueryRow(u.queries["HasUserRight"], userRight, userEntity.ID()).Scan(&hasUserRight)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Trace(err)
	}

	return hasUserRight == 1, nil
}

// Name returns the presentable name.
func (u *UsersStore) Name(user UserEntity, languageID uint16) string {
	var name string
	if user.FirstName() != nil {
		name = *user.FirstName()
	}
	if user.Surname() != nil {
		name += " " + *user.Surname()
	}
	if name == "" {
		name = u.translationsRepository.Singular(languageID, "user") + " " + user.ID()
	}
	return name
}

// NameWithEmail returns the presentable name with e-mail.
func (u *UsersStore) NameWithEmail(user UserEntity, languageID uint16) string {
	return u.Name(user, languageID) + " (" + user.Email() + ")"
}
