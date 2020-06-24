package fixtures

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const taxGroupQuery = `INSERT INTO "TaxGroup"("createdByID","active","code") VALUES($1,$2,$3) RETURNING "id"`
const taxGroupTranslationQuery = `INSERT INTO "TaxGroupTranslation"("createdByID","taxGroupID","language","field","value") VALUES($1,$2,$3,$4,$5)`
const taxQuery = `INSERT INTO "Tax"("createdByID","taxGroupID","country","rate") VALUES($1,$2,$3,$4) RETURNING "id"`

func (fixtures *Fixtures) taxes() error {
	// TaxGroup
	row := fixtures.inserterDatabase.QueryRow(taxGroupQuery, fixtures.mainUserID, true, "Hi")
	if err := row.Scan(&fixtures.taxGroupID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(taxGroupTranslationQuery, fixtures.mainUserID, fixtures.taxGroupID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "High"); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(taxGroupTranslationQuery, fixtures.mainUserID, fixtures.taxGroupID, fixtures.dutchLanguage.ID(), database.DBTranslationFieldName, "Hoog"); err != nil {
		return errors.Trace(err)
	}

	// Tax
	if _, err := fixtures.inserterDatabase.Exec(taxQuery, fixtures.mainUserID, fixtures.taxGroupID, fixtures.unitedKingdomCountry.ID(), 21.); err != nil {
		return errors.Trace(err)
	}

	return nil
}
