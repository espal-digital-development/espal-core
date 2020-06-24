package fixtures

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const propertyQuery = `INSERT INTO "Property"("createdByID","active","_type") VALUES($1,$2,$3) RETURNING "id"`
const propertyMultilingualQuery = `INSERT INTO "Property"("createdByID","active","_type","multiLingual") VALUES($1,$2,$3,$4) RETURNING "id"`
const propertyTranslationQuery = `INSERT INTO "PropertyTranslation"("createdByID","propertyID","language","field","value") VALUES($1,$2,$3,$4,$5)`
const propertyOptionQuery = `INSERT INTO "PropertyOption"("createdByID","propertyID","active") VALUES($1,$2,$3) RETURNING "id"`
const propertyOptionTranslationQuery = `INSERT INTO "PropertyOptionTranslation"("createdByID","propertyOptionID","language","field","value") VALUES($1,$2,$3,$4,$5)`

func (fixtures *Fixtures) properties() error {
	if err := fixtures.propertyName(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.propertyDescription(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.propertyImage(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.propertyColor(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.propertyLengthSize(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.propertyWidthSize(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.propertyPrice(); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyName() error {
	row := fixtures.inserterDatabase.QueryRow(propertyMultilingualQuery, fixtures.mainUserID, true, database.PropertytypeText, true)
	if err := row.Scan(&fixtures.propertyNameID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyNameID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Name"); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyNameID, fixtures.englishLanguage.ID(), database.DBTranslationFieldDescription, "Generic naming definition of a product"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyDescription() error {
	row := fixtures.inserterDatabase.QueryRow(propertyMultilingualQuery, fixtures.mainUserID, true, database.PropertytypeText, true)
	if err := row.Scan(&fixtures.propertyDescriptionID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyDescriptionID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Description"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyImage() error {
	row := fixtures.inserterDatabase.QueryRow(propertyQuery, fixtures.mainUserID, true, database.PropertytypeSinglefile)
	if err := row.Scan(&fixtures.propertyImageID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyImageID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Image"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyColor() error {
	row := fixtures.inserterDatabase.QueryRow(propertyQuery, fixtures.mainUserID, true, database.PropertytypeSingleselect)
	if err := row.Scan(&fixtures.propertyColorID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyColorID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Color"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionRedID string
	row = fixtures.inserterDatabase.QueryRow(propertyOptionQuery, fixtures.mainUserID, fixtures.propertyColorID, true)
	if err := row.Scan(&propertyOptionRedID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyOptionTranslationQuery, fixtures.mainUserID, propertyOptionRedID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Red"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionBlueID string
	row = fixtures.inserterDatabase.QueryRow(propertyOptionQuery, fixtures.mainUserID, fixtures.propertyColorID, true)
	if err := row.Scan(&propertyOptionBlueID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyOptionTranslationQuery, fixtures.mainUserID, propertyOptionBlueID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Blue"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyLengthSize() error {
	row := fixtures.inserterDatabase.QueryRow(propertyQuery, fixtures.mainUserID, true, database.PropertytypeSingleselect)
	if err := row.Scan(&fixtures.propertyLengthSizeID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyLengthSizeID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Length Size"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionSize34ID string
	row = fixtures.inserterDatabase.QueryRow(propertyOptionQuery, fixtures.mainUserID, fixtures.propertyLengthSizeID, true)
	if err := row.Scan(&propertyOptionSize34ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyOptionTranslationQuery, fixtures.mainUserID, propertyOptionSize34ID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "34"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionSize36ID string
	row = fixtures.inserterDatabase.QueryRow(propertyOptionQuery, fixtures.mainUserID, fixtures.propertyLengthSizeID, true)
	if err := row.Scan(&propertyOptionSize36ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyOptionTranslationQuery, fixtures.mainUserID, propertyOptionSize36ID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "36"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyWidthSize() error {
	row := fixtures.inserterDatabase.QueryRow(propertyQuery, fixtures.mainUserID, true, database.PropertytypeSingleselect)
	if err := row.Scan(&fixtures.propertyLengthWidthSizeID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyLengthWidthSizeID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Width Size"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionWidthSize36ID string
	row = fixtures.inserterDatabase.QueryRow(propertyOptionQuery, fixtures.mainUserID, fixtures.propertyLengthWidthSizeID, true)
	if err := row.Scan(&propertyOptionWidthSize36ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyOptionTranslationQuery, fixtures.mainUserID, propertyOptionWidthSize36ID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "36"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionWidthSize38ID string
	row = fixtures.inserterDatabase.QueryRow(propertyOptionQuery, fixtures.mainUserID, fixtures.propertyLengthWidthSizeID, true)
	if err := row.Scan(&propertyOptionWidthSize38ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyOptionTranslationQuery, fixtures.mainUserID, propertyOptionWidthSize38ID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "38"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (fixtures *Fixtures) propertyPrice() error {
	row := fixtures.inserterDatabase.QueryRow(propertyQuery, fixtures.mainUserID, true, database.PropertytypeCurrency)
	if err := row.Scan(&fixtures.propertyPriceID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyTranslationQuery, fixtures.mainUserID, fixtures.propertyPriceID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Price"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
