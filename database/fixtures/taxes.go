package fixtures

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const taxGroupQuery = `INSERT INTO "TaxGroup"("createdByID","active","code") VALUES($1,$2,$3) RETURNING "id"`
const taxGroupTranslationQuery = `INSERT INTO "TaxGroupTranslation"("createdByID","taxGroupID","language","field",
	"value") VALUES($1,$2,$3,$4,$5)`
const taxQuery = `INSERT INTO "Tax"("createdByID","taxGroupID","country","rate") VALUES($1,$2,$3,$4) RETURNING "id"`

func (f *Fixtures) taxes() error {
	// TaxGroup
	row := f.inserterDatabase.QueryRow(taxGroupQuery, f.mainUserID, true, "Hi")
	if err := row.Scan(&f.taxGroupID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(taxGroupTranslationQuery, f.mainUserID, f.taxGroupID, f.englishLanguage.ID(),
		database.DBTranslationFieldName, "High"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(taxGroupTranslationQuery, f.mainUserID, f.taxGroupID, f.dutchLanguage.ID(),
		database.DBTranslationFieldName, "Hoog"); err != nil {
		return errors.Trace(err)
	}

	// Tax
	if _, err := f.inserterDatabase.Exec(taxQuery, f.mainUserID, f.taxGroupID, f.unitedKingdomCountry.ID(),
		21.); err != nil {
		return errors.Trace(err)
	}

	return nil
}
