package fixtures

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const userGroupQuery = `INSERT INTO "UserGroup"("createdByID","active","userRights","currencies") VALUES($1,$2,$3,$4) RETURNING "id"`
const userGroupTranslationQuery = `INSERT INTO "UserGroupTranslation"("createdByID","userGroupID","language","field","value") VALUES($1,$2,$3,$4,$5)`

func (f *Fixtures) usersAndUserGroups() error {
	userGroupIDs := make([]string, 0)

	// Usergroup 1
	var userGroup1ID string
	row := f.inserterDatabase.QueryRow(userGroupQuery, f.mainUserID, true, f.userRightsBuffer.String(), "")
	if err := row.Scan(&userGroup1ID); err != nil {
		return errors.Trace(err)
	}
	userGroupIDs = append(userGroupIDs, userGroup1ID)
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup1ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Administrators"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup1ID, f.dutchLanguage.ID(), database.DBTranslationFieldName, "Beheerders"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup1ID, f.englishLanguage.ID(), database.DBTranslationFieldDescription, "Can manage all system functionalities"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup1ID, f.dutchLanguage.ID(), database.DBTranslationFieldDescription, "Kunnen alle systeem functionaliteiten beheren"); err != nil {
		return errors.Trace(err)
	}

	// Usergroup 2
	var userGroup2ID string
	row = f.inserterDatabase.QueryRow(userGroupQuery, f.mainUserID, true, f.userRightsBuffer.String(), "")
	if err := row.Scan(&userGroup2ID); err != nil {
		return errors.Trace(err)
	}
	userGroupIDs = append(userGroupIDs, userGroup2ID)
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup2ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Customers"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup2ID, f.dutchLanguage.ID(), database.DBTranslationFieldName, "Klanten"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup2ID, f.englishLanguage.ID(), database.DBTranslationFieldDescription, "Can order items and access shop functionality"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup2ID, f.dutchLanguage.ID(), database.DBTranslationFieldDescription, "Kunnen producten bestellen en shopfuncties gebruiken"); err != nil {
		return errors.Trace(err)
	}

	// Usergroup 3
	var userGroup3ID string
	row = f.inserterDatabase.QueryRow(userGroupQuery, f.mainUserID, true, f.userRightsBuffer.String(), "")
	if err := row.Scan(&userGroup3ID); err != nil {
		return errors.Trace(err)
	}
	userGroupIDs = append(userGroupIDs, userGroup3ID)
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup3ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "B2B Customers"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup3ID, f.dutchLanguage.ID(), database.DBTranslationFieldName, "B2B Klanten"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup3ID, f.englishLanguage.ID(), database.DBTranslationFieldDescription, "Can order items and access B2B shop functionality"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup3ID, f.dutchLanguage.ID(), database.DBTranslationFieldDescription, "Kunnen producten bestellen en B2B shopfuncties gebruiken"); err != nil {
		return errors.Trace(err)
	}

	// Usergroup 4
	var userGroup4ID string
	row = f.inserterDatabase.QueryRow(userGroupQuery, f.mainUserID, true, f.userRightsBuffer.String(), "")
	if err := row.Scan(&userGroup4ID); err != nil {
		return errors.Trace(err)
	}
	userGroupIDs = append(userGroupIDs, userGroup4ID)
	if _, err := f.inserterDatabase.Exec(userGroupQuery, f.mainUserID, true, "", ""); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup4ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Sales Agents"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup4ID, f.dutchLanguage.ID(), database.DBTranslationFieldName, "Vertegenwoordigers"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup4ID, f.englishLanguage.ID(), database.DBTranslationFieldDescription, "Can impersonate customers to trade for them"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(userGroupTranslationQuery, f.mainUserID, userGroup4ID, f.dutchLanguage.ID(), database.DBTranslationFieldDescription, "Kunnen klanten vertegenwoordigen door voor hun te handelen"); err != nil {
		return errors.Trace(err)
	}

	// User-Usergroup relationships
	userIDs := []string{f.mainUserID}
	for k := range userGroupIDs {
		for k2 := range userIDs {
			if _, err := f.inserterDatabase.Exec(`INSERT INTO "UserGroupUser"("createdByID","userGroupID","userID") VALUES($1,$2,$3)`, f.mainUserID, userGroupIDs[k], userIDs[k2]); err != nil {
				return errors.Trace(err)
			}
		}
	}

	return nil
}
